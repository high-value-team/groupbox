package backend

import (
	"net/http"
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
	return match("^/api/boxes$", reader.URL.Path)
}

func (handler *CreateBoxRequestHandler) Handle(writer http.ResponseWriter, reader *http.Request) {
	writeJsonResponse(writer, CreateBoxDTO{BoxKey: "1"})
}
