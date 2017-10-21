package interactions

import (
	"reflect"
	"testing"

	"github.com/high-value-team/groupbox/src/backend/models"
)

func TestMapToBoxDTO(t *testing.T) {
	// arrange
	boxKey := "1"
	box := models.Box{
		Title:        "Klassiker der Weltliteratur",
		CreationDate: "2017-10-01T10:30:59Z",
		Members: []models.Member{
			{
				Email:    "peter@acme.com",
				Nickname: "Golden Panda",
				Owner:    true,
				Key:      boxKey,
			},
		},
		Items: []models.Item{
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
	expected := &models.BoxDTO{
		Title:          "Klassiker der Weltliteratur",
		MemberNickname: "Golden Panda",
		CreationDate:   "2017-10-01T10:30:59Z",
		Items: []models.ItemDTO{
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
