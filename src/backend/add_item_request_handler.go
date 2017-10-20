package backend

import (
	"net/http"
	"regexp"
)

type AddItemRequestHandler struct {
	Interactions *Interactions
}

type AddItemRequestDTO struct {
	Message string `json:"message"`
}

func (handler *AddItemRequestHandler) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	boxKey, match := handler.Match(reader)
	if match {
		handler.Handle(writer, reader, boxKey)
		return true
	}
	return false
}

func (handler *AddItemRequestHandler) Match(reader *http.Request) (string, bool) {
	path := regexp.MustCompile("^/api/boxes/([a-zA-Z0-9]+)/items$")
	pathRegex := path.FindStringSubmatch(reader.URL.Path)
	if pathRegex != nil {
		return pathRegex[1], true
	}
	return "", false
}

func (handler *AddItemRequestHandler) Handle(writer http.ResponseWriter, reader *http.Request, boxKey string) {
	requestDTO := AddItemRequestDTO{}
	parseRequestBody(reader, &requestDTO)
	handler.Interactions.AddItem(boxKey, requestDTO.Message)
}
