package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//change this variables for your db
var (
	usr      = "usr"
	pwd      = "pwd"
	host     = "localhost"
	port     = 27017
	database = "database"
)

func GetCollection(collection string) *mongo.Collection {
	uri := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", usr, pwd, host, port, database)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
