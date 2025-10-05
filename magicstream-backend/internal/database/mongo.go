package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(ctx context.Context, uri, dbName string) (*Mongo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil { return nil, err }
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second); defer cancel()
	if err := client.Connect(ctx); err != nil { return nil, err }
	pingCtx, cancelPing := context.WithTimeout(context.Background(), 5*time.Second); defer cancelPing()
	if err := client.Ping(pingCtx, nil); err != nil { return nil, err }
	log.Println("connected to MongoDB")
	return &Mongo{Client: client, DB: client.Database(dbName)}, nil
}

func (m *Mongo) Close(ctx context.Context) error { return m.Client.Disconnect(ctx) }
