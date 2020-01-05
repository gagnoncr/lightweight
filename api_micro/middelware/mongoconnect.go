package middleware

import (
	"bytes"
	"context"
	"log"
	"os/exec"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb+srv://gagnonc:Xcv7Bnm!@endor-v1kqu.mongodb.net/test?retryWrites=true&w=majority"
	dbName           = "deployments"
	collName         = "deploy"
)

var (
	collection *mongo.Collection
)

func init() {

	cmd := exec.Command("/bin/sh", "-c", "cd middelware && ./mongo.sh")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(outb.String())
	log.Println(errb.String())

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
