package request_handlers

import (
	"net/http"

	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
)

type CreateBox struct {
	Interactions *interactions.Interactions
}

func NewCreateBoxHandler(interactions *interactions.Interactions) http.HandlerFunc {
	createBox := CreateBox{Interactions: interactions}
	return createBox.Handle
}

func (handler *CreateBox) Handle(writer http.ResponseWriter, reader *http.Request) {
	requestBody := JSONRequestCreateBox{}
	parseRequestBody(reader, &requestBody)
	owner := handler.Interactions.CreateBox(requestBody.Title, requestBody.Owner, requestBody.Members)
	writeJsonResponse(writer, JSONResponseCreateBox{BoxKey: owner.Key})
}
