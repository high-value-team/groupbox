package request_handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/backend/src/interior_interactions"
)

type AddItem struct {
	Interactions *interior_interactions.Interactions
}

func NewAddItemHandler(interactions *interior_interactions.Interactions) http.HandlerFunc {
	addItem := AddItem{Interactions: interactions}
	return addItem.Handle
}

func (handler *AddItem) Handle(writer http.ResponseWriter, reader *http.Request) {
	boxKey := chi.URLParam(reader, "boxKey")
	jsonRequest := JSONRequestAddItem{}
	parseRequestBody(reader, &jsonRequest)
	handler.Interactions.AddItem(boxKey, jsonRequest.Message)
}
