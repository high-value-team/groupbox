package request_handlers

import (
	"net/http"
	"regexp"
)

type GetBoxRequestHandler struct {
	Interactions *Interactions
}

func (handler *GetBoxRequestHandler) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	boxKey, match := handler.Match(reader)
	if match {
		handler.Handle(writer, reader, boxKey)
		return true
	}
	return false
}

func (handler *GetBoxRequestHandler) Match(reader *http.Request) (string, bool) {
	path := regexp.MustCompile("^/api/boxes/([a-zA-Z0-9]+)$")
	pathRegex := path.FindStringSubmatch(reader.URL.Path)
	if pathRegex != nil {
		return pathRegex[1], true
	}
	return "", false
}

func (handler *GetBoxRequestHandler) Handle(writer http.ResponseWriter, reader *http.Request, boxKey string) {
	boxDTO := handler.Interactions.GetBox(boxKey)
	writeJsonResponse(writer, boxDTO)
}
