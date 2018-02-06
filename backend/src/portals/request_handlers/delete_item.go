package request_handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/backend/src/interior/interactions"
)

type DeleteItem struct {
	Interactions *interactions.Interactions
}

func NewDeleteItemHandler(interactions *interactions.Interactions) http.HandlerFunc {
	updateItem := DeleteItem{Interactions: interactions}
	return updateItem.Handle
}

func (handler *DeleteItem) Handle(writer http.ResponseWriter, reader *http.Request) {
	boxKey := chi.URLParam(reader, "boxKey")
	itemID := chi.URLParam(reader, "itemID")
	handler.Interactions.DeleteItem(boxKey, itemID)
}
