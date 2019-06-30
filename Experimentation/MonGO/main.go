// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // You will be using this Trainer type later in the program
// type Trainer struct {
// 	Name string
// 	Age  int
// 	City string
// }

// func main() {
// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://cluster0-ydqrv.mongodb.net")

// 	// Connect to MongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	ash := Trainer{"Ash", 10, "Pallet Town"}

// 	collection := client.Database("test").Collection("trainers")
// 	collection.InsertOne(context.TODO(), ash)
// }

package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var username = "Carl"
var host1 = "cluster0-ydqrv.mongodb.net" // of the form foo.mongodb.net

func main() {
	ctx := context.TODO()
	pw, ok := os.LookupEnv("mongo_psw")

	if !ok {
		fmt.Println("error: unable to find MONGO_PW in the environment")
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
}
