package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Todo struct {
		Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Message   string        `json:"message"`
		CreatedOn time.Time     `json:"createdon,omitempty"`
	}
)
