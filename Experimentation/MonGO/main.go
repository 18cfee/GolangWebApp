package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var username string
var host1 string // of the form foo.mongodb.net

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	ctx := context.TODO()
	pw, ok := os.LookupEnv("mongo_psw")
	username, ok = os.LookupEnv("mongoUserName")
	host1, ok = os.LookupEnv("host1")

	if !ok {
		fmt.Println("error: unable to find in the environment")
		os.Exit(1)
	}
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", username, pw, host1)
	fmt.Println("connection string is:", mongoURI)

	// Set client options and connect
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("trainers")

	ash := Trainer{"Ash", 10, "Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	//brock := Trainer{"Brock", 15, "Pewter City"}

	for i := 0; i < 20; i++ {
		//insertResult, err := collection.InsertOne(context.TODO(), ash)
		collection.InsertOne(context.TODO(), ash)
		if err != nil {
			log.Fatal(err)
		}
	}

	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
