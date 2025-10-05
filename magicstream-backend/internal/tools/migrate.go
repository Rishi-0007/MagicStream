package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rishi-0007/magicstream-backend/internal/config"
	"github.com/rishi-0007/magicstream-backend/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	mongoConn, err := database.Connect(ctx, cfg.MongoURI, cfg.MongoDB)
	if err != nil { log.Fatal(err) }
	defer mongoConn.Close(ctx)
	col := mongoConn.DB.Collection("movies")

	pipe := mongo.Pipeline{
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "overview", Value: bson.D{{Key: "$ifNull", Value: bson.A{"$overview", "$admin_review"}}}},
			{Key: "genres", Value: bson.D{{Key: "$cond", Value: bson.D{
				{Key: "if", Value: bson.D{{Key: "$ne", Value: bson.A{bson.D{{Key: "$type", Value: "$genres"}}, "missing"}}}},
				{Key: "then", Value: "$genres"},
				{Key: "else", Value: bson.D{{Key: "$map", Value: bson.D{{Key: "input", Value: "$genre"}, {Key: "as", Value: "g"}, {Key: "in", Value: "$$g.genre_name"}}}}},
			}}}},
			{Key: "ranking", Value: bson.D{{Key: "$toDouble", Value: bson.D{{Key: "$ifNull", Value: bson.A{"$ranking.ranking_value", "$ranking.value", "$ranking"}}}}}},
		}}},
	}
	res, err := col.UpdateMany(ctx, bson.M{}, pipe)
	if err != nil { log.Fatal(err) }
	fmt.Printf("migrated %d movie docs to canonical schema\n", res.ModifiedCount)
}
