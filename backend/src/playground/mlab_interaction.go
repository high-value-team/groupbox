package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

// usage:
// export DB_USERNAME="groupbox"
// export DB_PASSWORD="geheim"
// export DB_NAME="develop"
// export COL_NAME="person"
// go run mlab_interaction.go
func main() {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COL_NAME")
	connectionString := fmt.Sprintf("mongodb://%s:%s@ds121565.mlab.com:21565/%s", username, password, databaseName)

	fmt.Printf("DB_USERNAME:%s\n", username)
	fmt.Printf("DB_PASSWORD:%s\n", password)
	fmt.Printf("DB_NAME:%s\n", databaseName)
	fmt.Printf("COL_NAME:%s\n", collectionName)
	fmt.Printf("connectionString:%s\n\n", connectionString)

	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB(databaseName).C(collectionName)

	person := Person{"Ale", "+55 53 8116 0002"}
	_, err = collection.Upsert(bson.M{"name": person.Name}, &person)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = collection.Find(bson.M{"name": person.Name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
