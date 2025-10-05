package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rishi-0007/magicstream-backend/internal/config"
	"github.com/rishi-0007/magicstream-backend/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	mongoConn, err := database.Connect(ctx, cfg.MongoURI, cfg.MongoDB)
	if err != nil { log.Fatal(err) }
	defer mongoConn.Close(ctx)
	createIndexes(ctx, mongoConn.DB)
}

func createIndexes(ctx context.Context, db *mongo.Database) {
	mov := db.Collection("movies")
	usr := db.Collection("users")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second); defer cancel()

	_, err := mov.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "imdb_id", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "title", Value: 1}}},
	})
	if err != nil { log.Fatal(err) }

	_, err = usr.Indexes().CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)})
	if err != nil { log.Fatal(err) }

	fmt.Println("indexes created")
}
