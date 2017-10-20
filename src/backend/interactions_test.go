package backend

import (
	"reflect"
	"testing"
)

func TestMapToBoxDTO(t *testing.T) {
	// arrange
	boxKey := "1"
	box := Box{
		Title:        "Klassiker der Weltliteratur",
		CreationDate: "2017-10-01T10:30:59Z",
		Members: []Member{
			{
				Email:    "peter@acme.com",
				Nickname: "Golden Panda",
				Owner:    true,
				Key:      boxKey,
			},
		},
		Items: []Item{
			{
				CreationDate: "2017-10-01T10:35:20Z",
				Subject:      "hallo",
				Message:      "Die drei Musketiere, Alexandre Dumas",
				AuthorKey:    boxKey,
			},
		},
	}
	interactions := Interactions{}

	// act
	actual := interactions.mapToBoxDTO(&box, boxKey)

	// assert
	expected := &BoxDTO{
		Title:          "Klassiker der Weltliteratur",
		MemberNickname: "Golden Panda",
		CreationDate:   "2017-10-01T10:30:59Z",
		Items: []ItemDTO{
			{
				AuthorNickname: "Golden Panda",
				CreationDate:   "2017-10-01T10:35:20Z",
				Subject:        "hallo",
				Message:        "Die drei Musketiere, Alexandre Dumas",
			},
		},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("mapToBoxDTO() failed!\nexpected:\n%+v\nactual:\n%+v\n", expected, actual)
	}
}


func TestExtractSubject(t *testing.T) {
	result := extractSubject("12345678901234567890")
	if result != "123456789012345..." {
		t.Errorf("Expected subject shorter than message; result: <%s>", result)
	}

	result = extractSubject("")
	if result != "?" {
		t.Errorf("Expected subject <?> for empty message", result)
	}

	result = extractSubject("1234567890")
	if result != "1234567890" {
		t.Errorf("Expected subject equal to message shorter than max subject len; result <%s>", result)
	}
}