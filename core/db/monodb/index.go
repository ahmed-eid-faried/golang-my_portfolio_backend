package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Index(collection *mongo.Collection) {
	// Create an index on the _id field (already unique by default)
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{Key: "_id", Value: 1}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
