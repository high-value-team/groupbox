package backend

import (
	"net/http"
	"regexp"
)

type CreateBoxRequestHandler struct {
	Interactions *Interactions
}

type CreateBoxDTO struct {
	BoxKey string `json:"boxKey"`
}

func (handler *CreateBoxRequestHandler) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	if handler.Match(reader) {
		handler.Handle(writer, reader)
		return true
	}
	return false
}

func (handler *CreateBoxRequestHandler) Match(reader *http.Request) bool {
	if regexp.MustCompile("^/api/boxes$") != nil {
		return true
	}
	return false
}

func (handler *CreateBoxRequestHandler) Handle(writer http.ResponseWriter, reader *http.Request) {
	writeJsonResponse(writer, CreateBoxDTO{BoxKey: "1"})
}
