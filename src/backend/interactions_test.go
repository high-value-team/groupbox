package backend

import (
	"reflect"
	"testing"
)

func TestMapToBoxDTO(t *testing.T) {
	// arrange
	box := Box{
		BoxID:        "litklas",
		Title:        "Klassiker der Weltliteratur",
		CreationDate: "2017-10-01T10:30:59Z",
		Members: []Member{
			{
				Email:    "peter@acme.com",
				Nickname: "Golden Panda",
				Owner:    true,
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
		},
	}
	boxMember := BoxMember{
		BoxKey: "1",
		BoxID:  "litklas",
		Member: Member{
			Email:    "peter@acme.com",
			Nickname: "Golden Panda",
			Owner:    true,
		},
	}
	interactions := Interactions{}

	// act
	actual := interactions.mapToBoxDTO(&box, &boxMember)

	// assert
	expected := &BoxDTO{
		MemberNickname: "Golden Panda",
		Title:          "Klassiker der Weltliteratur",
		CreationDate:   "2017-10-01T10:30:59Z",
		Items: []ItemDTO{
			{
				AuthorNickname: "Golden Panda",
				CreationDate:   "2017-10-01T10:35:20Z",
				Message:        "Die drei Musketiere, Alexandre Dumas",
			},
		},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("mapToBoxDTO() failed!\nexpected:\n%+v\nactual:\n%+v\n", expected, actual)
	}
}
