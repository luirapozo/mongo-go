package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meguca struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `json:"name"`
	Element string             `json:"element"`
	Atk     int                `json:"atk"`
	Def     int                `json:"def"`
	Hp      int                `json:"hp"`
	Rarity  int                `json:"rarity"`
	Magia   string             `json:"magia"`
	Doppel  string             `json:"doppel,omitempty"`
	Connect string             `json:"connect"`
	Ability string             `json:"ability"`
	Speed   int                `json:"speed"`
}

type Megucas []*Meguca
