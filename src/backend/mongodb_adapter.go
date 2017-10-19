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

func (adapter *MongoDBAdapter) openBox(err *error, boxKey string) *BoxMember {
	if *err != nil {
		return nil
	}

	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxMemberCollection)

	var boxMember BoxMember
	dbErr := collection.Find(bson.M{"boxKey": boxKey}).One(&boxMember)
	if checkError(err, dbErr) {
		return nil
	}

	return &boxMember
}

func (adapter *MongoDBAdapter) loadBox(err *error, boxID string) *Box {
	if *err != nil {
		return nil
	}

	sessionCopy := adapter.session.Copy()
	defer sessionCopy.Close()

	collection := sessionCopy.DB("").C(ConstBoxCollection)

	var box Box
	dbErr := collection.Find(bson.M{"boxId": boxID}).One(&box)
	if checkError(err, dbErr) {
		return nil
	}

	return &box
}

func checkError(err1 *error, err2 error) bool {
	if err2 != nil {
		if err2 == mgo.ErrNotFound {
			*err1 = SadError
		} else {
			*err1 = err2
		}
		return true
	}
	return false
}
