package dao

import (
	. "SmashLocalFinder/models"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LocalsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "locals"
)

func (l *LocalsDAO) Connect() {
	session, err := mgo.Dial(l.Server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	db = session.DB(l.Database)
}

func (l *LocalsDAO) FindAll() ([]Local, error) {
	var locals []Local
	err := db.C(COLLECTION).Find(bson.M{}).All(&locals)
	return locals, err
}
func (l *LocalsDAO) FindById(id string) (Local, error) {
	var local Local
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&local)
	return local, err
}

//Insert inserts new database object
func (l *LocalsDAO) Insert(local Local) error {
	fmt.Println(local)
	err := db.C(COLLECTION).Insert(&local)
	return err
}

func (l *LocalsDAO) Delete(local Local) error {
	err := db.C(COLLECTION).Remove(&local)
	return err
}

func (l *LocalsDAO) Update(local Local) error {
	err := db.C(COLLECTION).UpdateId(local.ID, &local)
	return err
}
