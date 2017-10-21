package providers

import (
	"time"

	"github.com/high-value-team/groupbox/src/backend/models"

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

func (adapter *MongoDBAdapter) LoadBox(boxKey string) *models.Box {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	var box models.Box
	err := collection.Find(bson.M{"members.key": boxKey}).One(&box)
	check(err)

	return &box
}
func (adapter *MongoDBAdapter) SaveBox(box *models.Box) {
	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	_, err := collection.Upsert(bson.M{"members.key": box.Members[0].Key}, box)
	check(err)
}

func check(err error) {
	if err != nil {
		if err == mgo.ErrNotFound {
			panic(models.SadException{Err: mgo.ErrNotFound})
		} else {
			panic(models.SuprisingException{Err: err})
		}
	}
}
