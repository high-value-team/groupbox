package request_handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
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
	boxDTO := handler.Interactions.GetBox(boxKey)
	writeJsonResponse(writer, boxDTO)
}
