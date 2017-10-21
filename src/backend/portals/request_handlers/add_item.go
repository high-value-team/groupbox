package request_handlers

import (
	"net/http"
	"regexp"

	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
	"github.com/high-value-team/groupbox/src/backend/models"
)

type AddItem struct {
	Interactions *interactions.Interactions
}

func (handler *AddItem) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	boxKey, match := handler.Match(reader)
	if match {
		handler.Handle(writer, reader, boxKey)
		return true
	}
	return false
}

func (handler *AddItem) Match(reader *http.Request) (string, bool) {
	path := regexp.MustCompile("^/api/boxes/([a-zA-Z0-9]+)/items$")
	pathRegex := path.FindStringSubmatch(reader.URL.Path)
	if pathRegex != nil {
		return pathRegex[1], true
	}
	return "", false
}

func (handler *AddItem) Handle(writer http.ResponseWriter, reader *http.Request, boxKey string) {
	requestDTO := models.AddItemRequestDTO{}
	parseRequestBody(reader, &requestDTO)
	handler.Interactions.AddItem(boxKey, requestDTO.Message)
}
