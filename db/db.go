package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri = os.Getenv("DB")

	Lobbies *mongo.Collection
)

func Init() (err error) {
	// Get options from the connection URI.
	opts := options.Client().ApplyURI(uri)

	// Create a client with the options.
	client, err := mongo.NewClient(opts)
	if err != nil {
		return
	}

	// Connect to the server.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	// Verify the client is connected with a ping.
	err = client.Ping(ctx, opts.ReadPreference)
	if err != nil {
		return
	}

	// Get the DB.
	db := client.Database("duopoly")

	// Get all collections from the DB.
	Lobbies = db.Collection("lobbies")
	return
}
