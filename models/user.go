package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `json:"name"`
	UserName    string             `json:"username"`
	Password    string             `json:"password"`
	Rank        uint               `json:"rank"`
	Time_create time.Time          `json:"time_create"`
	Characters  Megucas            `json:"characters,omitempty"`
}

type Users []*User
