package backend

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	mgo "gopkg.in/mgo.v2"
)

var ConnectionString string = ""

// export USERNAME="groupbox"
// export PASSWORD="geheim"
// export DB_NAME="develop"
// go test
func TestMain(m *testing.M) {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	ConnectionString = fmt.Sprintf("mongodb://%s:%s@ds121565.mlab.com:21565/%s", username, password, databaseName)
	os.Exit(m.Run())
}

func TestOpenBoxFound(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	var err error
	actual := adapter.openBox(&err, "1")

	// assert
	if err != nil {
		t.Fatal(err)
	}

	expected := &BoxMember{
		BoxKey: "1",
		BoxID:  "litklas",
		Member: Member{
			Email:    "peter@acme.com",
			Nickname: "Golden Panda",
			Owner:    true,
		},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("openBox() failed!\nexpected:\n%+v\nactual:\n%+v\n", expected, actual)
	}
}

func TestOpenBoxNOTFound(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	var err error
	actual := adapter.openBox(&err, "key does not exist")

	// assert
	if err != mgo.ErrNotFound {
		t.Fatal("expected ErrNotFound error")
	}
	if actual != nil {
		t.Fatal("expected boxMember to be nil")
	}
}

func TestOpenBoxErrorPassedIn(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	err := fmt.Errorf("my error")
	expectedErr := err
	actual := adapter.openBox(&err, "1")

	// assert
	if err != expectedErr {
		t.Fatal("expected expectedErr error")
	}
	if actual != nil {
		t.Fatal("expected boxMember to be nil")
	}
}

func TestLoadBoxFound(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	var err error
	actual := adapter.loadBox(&err, "litklas")

	// assert
	if err != nil {
		t.Fatal(err)
	}

	expected := &Box{
		BoxID:        "litklas",
		Title:        "Klassiker der Weltliteratur",
		CreationDate: "2017-10-01T10:30:59Z",
		Members: []Member{
			{
				Email:    "peter@acme.com",
				Nickname: "Golden Panda",
				Owner:    true,
			},
			{
				Email:    "paul@acme.com",
				Nickname: "Flying Fox",
				Owner:    false,
			},
			{
				Email:    "mary@acme.com",
				Nickname: "Fierce Tiger",
				Owner:    false,
			},
		},
		Items: []Item{
			{
				CreationDate: "2017-10-01T10:35:20Z",
				Message:      "Die drei Musketiere, Alexandre Dumas",
				Author: Member{
					Email:    "peter@acme.com",
					Nickname: "Golden Panda",
					Owner:    true,
				},
			},
			{
				CreationDate: "2017-10-02T14:40:30Z",
				Message:      "Der Zauberer von Oz, Frank Baum",
				Author: Member{
					Email:    "mary@acme.com",
					Nickname: "Fierce Tiger",
					Owner:    false,
				},
			},
			{
				CreationDate: "2017-10-03T20:55:10Z",
				Message:      "Schuld und Sühne, Dostojewski, www.amazon.de/Schuld-Sühne-Fjodr-Michailowitsch-Dostojewski-ebook/dp/B004UBCWK6",
				Author: Member{
					Email:    "paul@acme.com",
					Nickname: "Flying Fox",
					Owner:    false,
				},
			},
		},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("loadBox() failed!\nexpected:\n%+v\nactual:\n%+v\n", expected, actual)
	}
}

func TestLoadBoxNOTFound(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	var err error
	actual := adapter.loadBox(&err, "key does not exist")

	// assert
	if err != mgo.ErrNotFound {
		t.Fatal("expected ErrNotFound error")
	}
	if actual != nil {
		t.Fatal("expected box to be nil")
	}
}

func TestLoadBoxErrorPassedIn(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	err := fmt.Errorf("my error")
	expectedErr := err
	actual := adapter.loadBox(&err, "litklas")

	// assert
	if err != expectedErr {
		t.Fatal("expected expectedErr error")
	}
	if actual != nil {
		t.Fatal("expected box to be nil")
	}
}
