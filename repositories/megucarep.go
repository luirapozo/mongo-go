package repositories

import (
	"mongo-go/database"
	"mongo-go/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var collectionM = database.GetCollection("meguca")

func ReadM() (models.Megucas, error) {
	filter := bson.M{}
	cur, err := collectionM.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var megucas models.Megucas
	for cur.Next(ctx) {
		var meguca models.Meguca
		err = cur.Decode(&meguca)
		if err != nil {
			return nil, err
		}
		megucas = append(megucas, &meguca)
	}
	return megucas, nil
}

func ReadAbilities(id string) (models.Meguca, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Meguca{}, err
	}
	filter := bson.M{"_id": objectId}
	result := collectionM.FindOne(ctx, filter)
	meguca := models.Meguca{}
	result.Decode(&meguca)

	return meguca, nil
}

func FindMegucaByName(name string) models.Meguca {

	filter := bson.M{"name": name}
	result := collectionM.FindOne(ctx, filter)
	meguca := models.Meguca{}
	result.Decode(&meguca)

	return meguca
}
