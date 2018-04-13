package request_handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/backend/src/exceptions"
	"github.com/high-value-team/groupbox/backend/src/interior/interactions"
	interiorModels "github.com/high-value-team/groupbox/backend/src/interior/models"
)

type GetBox struct {
	Interactions *interactions.Interactions
}

func NewGetBoxHandler(interactions *interactions.Interactions) http.HandlerFunc {
	getBox := GetBox{Interactions: interactions}
	return getBox.Handle
}

func (handler *GetBox) Handle(writer http.ResponseWriter, reader *http.Request) {
	boxKey := chi.URLParam(reader, "boxKey")
	box := handler.Interactions.GetBox(boxKey)
	jsonBox := mapToJSONBox(box, boxKey)
	writeJsonResponse(writer, jsonBox)
}

func mapToJSONBox(box *interiorModels.Box, boxKey string) *JSONBox {
	requestingMember := selectMember(boxKey, box.Members)
	jsonBox := JSONBox{
		Title:          box.Title,
		MemberNickname: requestingMember.Nickname,
		CreationDate:   box.CreationDate.Format(time.RFC3339),
		Items:          []JSONItem{},
	}
	for i := range box.Items {
		jsonBox.Items = append(jsonBox.Items, JSONItem{
			ItemID:         strconv.Itoa(i),
			AuthorNickname: selectMember(box.Items[i].AuthorKey, box.Members).Nickname,
			CreationDate:   box.Items[i].CreationDate.Format(time.RFC3339),
			Subject:        box.Items[i].Subject,
			Message:        box.Items[i].Message,
		})
	}
	return &jsonBox
}

func selectMember(key string, members []interiorModels.Member) *interiorModels.Member {
	for i := range members {
		if members[i].Key == key {
			return &members[i]
		}
	}
	panic(exceptions.SuprisingException{Err: fmt.Errorf("No member found for key:%s!", key)})
}
