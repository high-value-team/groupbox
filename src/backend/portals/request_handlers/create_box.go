package request_handlers

import (
	"net/http"

	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
	"github.com/high-value-team/groupbox/src/backend/models"
)

type CreateBox struct {
	Interactions *interactions.Interactions
}

func (handler *CreateBox) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	if handler.Match(reader) {
		handler.Handle(writer, reader)
		return true
	}
	return false
}

func (handler *CreateBox) Match(reader *http.Request) bool {
	return match("^/api/boxes$", reader.URL.Path)
}

func (handler *CreateBox) Handle(writer http.ResponseWriter, reader *http.Request) {
	requestDTO := models.CreateBoxRequestDTO{}
	parseRequestBody(reader, &requestDTO)
	responseDTO := handler.Interactions.CreateBox(requestDTO.Title, requestDTO.Owner, requestDTO.Members)
	writeJsonResponse(writer, responseDTO)
}
