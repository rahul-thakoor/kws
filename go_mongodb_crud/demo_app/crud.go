package main

import (
	"log"
	"context"
	"time"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	// Create context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Create client
	client, err := mongo.Connect(ctx, "mongodb://supinfo:supinfo123@ds213645.mlab.com:13645/supinfo-kws-demo")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection successful")


	// create a Module type representing a particular course module
	type Module struct{
		Id string
		Title string
		ECTS int
		Description string
	}

	// Get collection
	collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")


	deleteResult, err := collection.DeleteOne(ctx, bson.D{{"id", "2WEB"}})

	collection := client.Database("supinfo-kws-demo").Collection("Assignments")
	insertResult, err := collection.InsertOne(ctx, bson.M{"ID":"MP1", "Description":"Puzzler- Use HTML5 to create a mini puzzle game", "Module":"1WEB" , "Deadline":"25-09-2018" })

	if err != nil {
		log.Fatal(err)
	}


	log.Println("No. of deleted document(s) : ", deleteResult.DeletedCount)

	// close connection
	client.Disconnect(ctx)


	log.Println("Inserted a single document: ", insertResult.InsertedID)
	// close connection
	client.Disconnect(ctx)

}