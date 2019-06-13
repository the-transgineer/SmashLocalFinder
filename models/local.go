package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Local Represents a local tournament
type Local struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Day       string        `bson:"day" json:"day"`
	Location  string        `bson:"location" json:"location"`
	Region    string        `bson:"region" json:"region"`
	Organizer string        `bson:"orgranizer" json:"organizer"`
	Time      time.Time     `bson:"time" json:"time"`
}
