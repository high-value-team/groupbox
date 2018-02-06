package providers

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/high-value-team/groupbox/backend/src/models"
)

var ConnectionString string = ""

// export MONGODB_URL=mongodb://<username>:<password>@ds121565.mlab.com:21565/<databasename>
// go test
func TestMain(m *testing.M) {
	ConnectionString = os.Getenv("MONGODB_URL")
	os.Exit(m.Run())
}

func TestSaveBox(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()
	var exception interface{}
	box := &models.Box{
		Title:        "Klassiker der Weltliteratur",
		CreationDate: time.Date(2017, 10, 1, 10, 30, 59, 0, time.UTC), // "2017-10-01T10:30:59Z",
		Members: []models.Member{
			{
				Key:      "1",
				Email:    "peter@acme.com",
				Nickname: "Golden Panda",
				Owner:    true,
			},
			{
				Key:      "2",
				Email:    "paul@acme.com",
				Nickname: "Flying Fox",
				Owner:    false,
			},
			{
				Key:      "3",
				Email:    "mary@acme.com",
				Nickname: "Fierce Tiger",
				Owner:    false,
			},
		},
		Items: []models.Item{
			{
				CreationDate: time.Date(2017, 10, 1, 10, 35, 20, 0, time.UTC),
				Subject:      "Die drei Muske...",
				Message:      "Die drei Musketiere, Alexandre Dumas",
				AuthorKey:    "1",
			},
			{
				CreationDate: time.Date(2017, 10, 2, 14, 40, 30, 0, time.UTC),
				Subject:      "Der Zauberer v...",
				Message:      "Der Zauberer von Oz, Frank Baum",
				AuthorKey:    "3",
			},
			{
				CreationDate: time.Date(2017, 10, 3, 20, 55, 10, 0, time.UTC),
				Subject:      "Schuld und Süh...",
				Message:      "Schuld und Sühne, Dostojewski, www.amazon.de/Schuld-Sühne-Fjodr-Michailowitsch-Dostojewski-ebook/dp/B004UBCWK6",
				AuthorKey:    "2",
			},
		},
	}

	// act
	recoverFromPanic(
		func() {
			adapter.SaveBox(box)
		},
		func(recovered interface{}) {
			exception = recovered
		},
	)

	// assert
	if _, ok := exception.(models.SadException); ok {
		t.Fatalf("did not expect exception: %v", exception)
	}
}

func TestLoadBoxFound(t *testing.T) {
	// arrange
	adapter := MongoDBAdapter{ConnectionString: ConnectionString}
	adapter.Start()
	defer adapter.Stop()

	// act
	var err error
	actual := adapter.LoadBox("1")

	// assert
	if err != nil {
		t.Fatal(err)
	}

	expected := &models.Box{
		Title:        "Klassiker der Weltliteratur",
		CreationDate: time.Date(2017, 10, 1, 10, 30, 59, 0, time.UTC), // "2017-10-01T10:30:59Z",
		Members: []models.Member{
			{
				Key:      "1",
				Email:    "peter@acme.com",
				Nickname: "Golden Panda",
				Owner:    true,
			},
			{
				Key:      "2",
				Email:    "paul@acme.com",
				Nickname: "Flying Fox",
				Owner:    false,
			},
			{
				Key:      "3",
				Email:    "mary@acme.com",
				Nickname: "Fierce Tiger",
				Owner:    false,
			},
		},
		Items: []models.Item{
			{
				CreationDate: time.Date(2017, 10, 1, 10, 35, 20, 0, time.UTC),
				Subject:      "Die drei Muske...",
				Message:      "Die drei Musketiere, Alexandre Dumas",
				AuthorKey:    "1",
			},
			{
				CreationDate: time.Date(2017, 10, 2, 14, 40, 30, 0, time.UTC),
				Subject:      "Der Zauberer v...",
				Message:      "Der Zauberer von Oz, Frank Baum",
				AuthorKey:    "3",
			},
			{
				CreationDate: time.Date(2017, 10, 3, 20, 55, 10, 0, time.UTC),
				Subject:      "Schuld und Süh...",
				Message:      "Schuld und Sühne, Dostojewski, www.amazon.de/Schuld-Sühne-Fjodr-Michailowitsch-Dostojewski-ebook/dp/B004UBCWK6",
				AuthorKey:    "2",
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
	var exception interface{}
	var actual interface{}
	recoverFromPanic(
		func() {
			actual = adapter.LoadBox("key does not exist")
		},
		func(recovered interface{}) {
			exception = recovered
		},
	)

	// assert
	if _, ok := exception.(models.SadException); !ok {
		t.Fatal("expected SadException")
	}

	if actual != nil {
		t.Fatal("expected actual to be nil")
	}
}

func recoverFromPanic(callback func(), onPanic func(r interface{})) {
	defer func() {
		r := recover()
		if r != nil {
			onPanic(r)
		}
	}()
	callback()
}
