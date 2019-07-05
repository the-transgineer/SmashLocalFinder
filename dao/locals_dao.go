package dao

import (
	"fmt"
	"log"

	. "github.com/the-transgineer/SmashLocalFinder/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//LocalsDAO is a Data Access Object
type LocalsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

//COLLECTION is a constant
const (
	COLLECTION = "locals"
)

//Connect connects to database
func (l *LocalsDAO) Connect() {
	session, err := mgo.Dial(l.Server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	db = session.DB(l.Database)
}

//FindAll returns all locals
func (l *LocalsDAO) FindAll() ([]Local, error) {
	var locals []Local
	err := db.C(COLLECTION).Find(bson.M{}).All(&locals)
	return locals, err
}

//FindById returns local based on ID
func (l *LocalsDAO) FindById(id string) (Local, error) {
	var local Local
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&local)
	return local, err
}

//Insert inserts new database object
func (l *LocalsDAO) Insert(local Local) error {
	err := db.C(COLLECTION).Insert(&local)
	return err
}

//Delete removes a local
func (l *LocalsDAO) Delete(local Local) error {
	err := db.C(COLLECTION).Remove(&local)
	return err
}

//Update Updates a local based on ID
func (l *LocalsDAO) Update(local Local) error {
	err := db.C(COLLECTION).UpdateId(local.ID, &local)
	return err
}

//FindByRegion returns all locals based on a region
func (l *LocalsDAO) FindByRegion(region string) ([]Local, error) {
	var locals []Local
	err := db.C(COLLECTION).Find(bson.M{"region": region}).All(&locals)
	return locals, err
}
