package db

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := "mongodb://mongo:27017" // Docker service name
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    log.Println("Connected to MongoDB")
    Client = client
}

func NotesCollection() *mongo.Collection {
    return Client.Database("keep").Collection("notes")
}