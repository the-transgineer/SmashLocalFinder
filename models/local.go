package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Local Represents a local tournament
type Local struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Day        string        `bson:"day" json:"day"`
	Location   string        `bson:"location" json:"location"`
	Region     string        `bson:"region" json:"region"`
	Organizers []string      `bson:"orgranizers" json:"organizers"`
}
