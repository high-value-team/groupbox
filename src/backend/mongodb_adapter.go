package backend

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ConstBoxCollection       string = "Box"
	ConstBoxMemberCollection string = "BoxMember"
)

type MongoDBAdapter struct {
	ConnectionString string
	session          *mgo.Session
}

func (adapter *MongoDBAdapter) Start() {
	var err error
	adapter.session, err = mgo.DialWithTimeout(adapter.ConnectionString, 1*time.Second)
	if err != nil {
		panic(err)
	}
}

func (adapter *MongoDBAdapter) Stop() {
	adapter.session.Close()
}

func (adapter *MongoDBAdapter) openBox(boxKey string) *BoxMember {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxMemberCollection)

	var boxMember BoxMember
	err := collection.Find(bson.M{"boxKey": boxKey}).One(&boxMember)
	check(err)

	return &boxMember
}

func (adapter *MongoDBAdapter) loadBox(boxID string) *Box {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	var box Box
	err := collection.Find(bson.M{"boxId": boxID}).One(&box)
	check(err)

	return &box
}

func check(err error) {
	if err != nil {
		if err == mgo.ErrNotFound {
			panic(SadException{Err: mgo.ErrNotFound})
		} else {
			panic(SuprisingException{Err: err})
		}
	}
}
