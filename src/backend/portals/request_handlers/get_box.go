package request_handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
	"github.com/high-value-team/groupbox/src/backend/models"
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

func mapToJSONBox(box *models.Box, boxKey string) *JSONBox {
	requestingMember := selectMember(boxKey, box.Members)
	jsonBox := JSONBox{
		Title:          box.Title,
		MemberNickname: requestingMember.Nickname,
		CreationDate:   box.CreationDate.Format(time.RFC3339),
		Items:          []JSONItem{},
	}
	for _, item := range box.Items {
		jsonBox.Items = append(jsonBox.Items, JSONItem{
			AuthorNickname: selectMember(item.AuthorKey, box.Members).Nickname,
			CreationDate:   item.CreationDate.Format(time.RFC3339),
			Subject:        item.Subject,
			Message:        item.Message,
		})
	}
	return &jsonBox
}

func selectMember(key string, members []models.Member) *models.Member {
	for i := range members {
		if members[i].Key == key {
			return &members[i]
		}
	}
	panic(models.SuprisingException{Err: fmt.Errorf("No member found for key:%s!", key)})
}
