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
type Counts struct {
	VarName string
	Val     int
	Rando   string
}

type Select struct {
	VarName string
}

type Customer struct {
	Val int
}

var client *mongo.Client

func InitMongo() {
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
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
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

// func CreateCustomer(id int) error {

// }

func GetNextCustId() (int, error) {
	collection := client.Database("cabin").Collection("counts")

	num := Customer{0}
	//ash := Trainer{"Carl", 10, "Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	//brock := Trainer{"Brock", 15, "Pewter City"}

	// for i := 0; i < 20; i++ {
	// 	//insertResult, err := collection.InsertOne(context.TODO(), ash)
	// 	collection.InsertOne(context.TODO(), ash)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// create a value into which the result can be decoded
	var result Counts

	err := collection.FindOne(context.TODO(), num).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result.Val, nil
}

func CloseMongo() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func main() {
	InitMongo()
	id, err := GetNextCustId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the id is ", id)
	CloseMongo()
}
