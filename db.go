package main

import (
	"context"
	"log"
    "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
    dbName = "users"
    URI = "mongodb://127.0.0.1:27017/" + dbName
)

func Connect() *mongo.Database {
    client, err := mongo.NewClient(options.Client().ApplyURI(URI))
    if err != nil {
        log.Panic(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 45)
    defer cancel()

    client.Connect(ctx)

    return client.Database(dbName)
}
