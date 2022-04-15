package repositories

import (
	"context"
	"mongo-go/database"
	"mongo-go/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var collectionU = database.GetCollection("users")
var ctx = context.Background()

func CreateU(user models.User) error {
	_, err := collectionU.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func ReadUsers() (models.Users, error) {
	filter := bson.M{}
	cur, err := collectionU.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var users models.Users
	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func FindUserByUName(username string, password string) models.User {
	filter := bson.M{"username": username, "password": password}
	result := collectionU.FindOne(ctx, filter)
	user := models.User{}
	result.Decode(&user)

	return user
}
func UpdateU(user models.User, oid primitive.ObjectID) error {
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"rank":       user.Rank,
			"characters": user.Characters,
		},
	}

	_, err := collectionU.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteU(userId string) error {
	oid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	_, err = collectionU.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
