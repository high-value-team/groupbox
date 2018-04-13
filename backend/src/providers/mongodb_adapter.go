package providers

import (
	"time"

	"github.com/high-value-team/groupbox/backend/src/exceptions"
	interiorModels "github.com/high-value-team/groupbox/backend/src/interior/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ConstBoxCollection string = "Box"
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

func (adapter *MongoDBAdapter) LoadBox(boxKey string) *interiorModels.Box {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	var bsonBox BSONBox
	err := collection.Find(bson.M{"members.key": boxKey}).One(&bsonBox)
	check(err)

	return ToBox(&bsonBox)
}

func (adapter *MongoDBAdapter) SaveBox(box *interiorModels.Box) {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	bsonBox := ToBSONBox(box)
	collection := sessionCopy.DB("").C(ConstBoxCollection)

	_, err := collection.Upsert(bson.M{"members.key": bsonBox.Members[0].Key}, bsonBox)
	check(err)
}

func check(err error) {
	if err != nil {
		if err == mgo.ErrNotFound {
			panic(exceptions.SadException{Err: mgo.ErrNotFound})
		} else {
			panic(exceptions.SuprisingException{Err: err})
		}
	}
}
