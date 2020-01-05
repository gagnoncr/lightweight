package middleware

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
)

const (
	dbName           = "deployments"
	collName         = "deploy"
)

var (
	collection *mongo.Collection
)

func init() {

	b, err := ioutil.ReadFile("middelware/mongoSensitive.encrypted")

	if err != nil {
		fmt.Print(err)
	}

	connectionString := string(b)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Connection to MongoDB was successful")

	collection = client.Database(dbName).Collection(collName)

	log.Printf("Collection created")
}
