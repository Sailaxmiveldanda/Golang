package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/sailaxmiveldanda/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/v2/mongo"
)

// for errors ot clear ctrl + shift + p
const connectionString = "mongodb+srv://veldandasailaxmi_db_user:Zje8OoJBJh0vZVaT@cluster0.fr9ze5q.mongodb.net/?appName=Cluster0"
const dbName = "netflix"
const colName = " watchlist"

// important connection
var collection *mongo.Collection

//connect with mongodb

func init() { //initilization method runs only 1 time
	//cliient options to give
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance(refrence)
	fmt.Println("Collection Reference is Ready")
}

//mongodb helpers

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie in db with id: ", inserted.InsertedID)
}
