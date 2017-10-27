package request_handlers

import (
	"net/http"

	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
	"github.com/high-value-team/groupbox/src/backend/models"
)

type CreateBox struct {
	Interactions *interactions.Interactions
}

func NewCreateBoxHandler(interactions *interactions.Interactions) http.HandlerFunc {
	createBox := CreateBox{Interactions: interactions}
	return createBox.Handle
}

func (handler *CreateBox) Handle(writer http.ResponseWriter, reader *http.Request) {
	requestDTO := models.CreateBoxRequestDTO{}
	parseRequestBody(reader, &requestDTO)
	responseDTO := handler.Interactions.CreateBox(requestDTO.Title, requestDTO.Owner, requestDTO.Members)
	writeJsonResponse(writer, responseDTO)
}
