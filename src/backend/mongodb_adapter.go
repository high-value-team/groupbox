package backend

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
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
	log.Printf("Connecting to MongoDB <%s>...", adapter.ConnectionString)

	var err error
	adapter.session, err = mgo.DialWithTimeout(adapter.ConnectionString, 1*time.Second)
	if err != nil {
		panic(err)
	}
}

func (adapter *MongoDBAdapter) Stop() {
	adapter.session.Close()
}

func (adapter *MongoDBAdapter) loadBox(boxKey string) *Box {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	var box Box
	err := collection.Find(bson.M{"members.key": boxKey}).One(&box)
	check(err)

	return &box
}
func (adapter *MongoDBAdapter) saveBox(box *Box) {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	err := collection.Insert(box)
	check(err)
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
