package backend

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"
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

func writeJsonResponse(writer http.ResponseWriter, i interface{}) {
	writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(i)
}
