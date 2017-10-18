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
