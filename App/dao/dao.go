package dao

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var username string
var host1 string // of the form foo.mongodb.net
var client *mongo.Client

func RetrieveCustomers() ([]Customer, error) {
	collection := client.Database("cabin").Collection("customers")
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"id": -1})

	// Here's an array in which you can store the decoded documents
	var results []Customer

	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Customer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}

func init() {
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

type Customer struct {
	Name  string
	Id    int
	Forms string
}

func GetHighestCustId() (int, error) {
	collection := client.Database("cabin").Collection("customers")
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"id": -1})
	findOptions.SetLimit(1)

	cur, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println("line 64 failure")
		return 0, err
	}

	var cust Customer

	// could iterate through results, but there is only one now
	cur.Next(context.TODO())
	err = cur.Decode(&cust)
	fmt.Println(cust.Id)
	if err != nil {
		return 0, err
	}

	err = cur.Err()

	cur.Close(context.TODO())

	return cust.Id, err
}

func CreateNewCust(id int) error {
	collection := client.Database("cabin").Collection("customers")
	newCust := Customer{Id: id}
	result, err := collection.InsertOne(context.TODO(), newCust)
	fmt.Println(result)
	return err
}

func CloseMongo() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
