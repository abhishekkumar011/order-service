//This file connects your Go application to MongoDB and stores the database collection so the rest of your app can use it.
package database

import (
	"context"
	"log"
	"time"

	"food-delivery-order/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var OrderCollection *mongo.Collection

func Connect(cfg config.Config) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(cfg.MongoURI),
	)

	if err != nil {
		log.Fatal(err)
	}

	Client = client

	OrderCollection = client.
		Database(cfg.DatabaseName).
		Collection(cfg.CollectionName)

	log.Println("MongoDB Connected")
}

/*
	1. Context - This is one of the biggest differences between Node and Go.
				 - Node - await User.find()
				 - Go   - Collection.Find(ctx)
	   - Everything in Go accepts a Context.
	   - Every MongoDB operation in Go needs a Context.
	2. Client stores the MongoDB connection.
	3. Go doesn't use try/catch
	4. ctx, cancel := context.WithTimeout(...) - Creates a context with a 10-second timeout. Try connecting to MongoDB, but if it takes more than 10 seconds, stop
	5. defer cancel() - This is Go Keyword Think of like Run this function when the current function finishes. Here it cleans up the context after Connect() ends.
	6. var Client *mongo.Client - Stores the MongoDB connection like const client = mongoose.connection in node.js app
	7. var OrderCollection *mongo.Collection - Stores a reference to the orders collection. like const Order = mongoose.model("Order", orderSchema); in express.js
*/
