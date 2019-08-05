package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Student struct {
		ID       bson.ObjectId `json:"id" bson:"_id"`
		Name     string        `json:"name" bson:"name"`
		Email    string        `json:"email" bson:"email"`
		Password string        `json:"password" bson:"password"`
		Gender   string        `json:"gender" bson:"gender"`
	}
)
