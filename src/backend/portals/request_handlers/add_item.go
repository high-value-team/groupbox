package request_handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
)

type AddItem struct {
	Interactions *interactions.Interactions
}

func NewAddItemHandler(interactions *interactions.Interactions) http.HandlerFunc {
	addItem := AddItem{Interactions: interactions}
	return addItem.Handle
}

func (handler *AddItem) Handle(writer http.ResponseWriter, reader *http.Request) {
	boxKey := chi.URLParam(reader, "boxKey")
	jsonRequest := JSONRequestAddItem{}
	parseRequestBody(reader, &jsonRequest)
	handler.Interactions.AddItem(boxKey, jsonRequest.Message)
}
