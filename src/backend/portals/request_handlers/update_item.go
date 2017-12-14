package request_handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
)

type UpdateItem struct {
	Interactions *interactions.Interactions
}

func NewUpdateItemHandler(interactions *interactions.Interactions) http.HandlerFunc {
	updateItem := UpdateItem{Interactions: interactions}
	return updateItem.Handle
}

func (handler *UpdateItem) Handle(writer http.ResponseWriter, reader *http.Request) {
	boxKey := chi.URLParam(reader, "boxKey")
	jsonRequest := JSONRequestUpdateItem{}
	parseRequestBody(reader, &jsonRequest)
	handler.Interactions.UpdateItem(boxKey, jsonRequest.Subject, jsonRequest.Message)
}
