package backend

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type GetBoxRequestHandler struct {
	Interactions *Interactions
}

func (handler *GetBoxRequestHandler) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	if handler.Match(reader) {
		handler.Handle(writer, reader)
		return true
	}
	return false
}

func (handler *GetBoxRequestHandler) Match(reader *http.Request) bool {
	path := regexp.MustCompile("^/api/boxes/([a-zA-Z0-9]+)$")
	return path.FindStringSubmatch(reader.URL.Path) != nil
}

func (handler *GetBoxRequestHandler) Handle(writer http.ResponseWriter, reader *http.Request) {
	boxKey := handler.parseBoxKey(reader.URL.Path)
	boxDTO, err := handler.Interactions.GetBox(boxKey)
	if err != nil {
		// TODO
		//switch err.Type {
		//case MySadError:
		//	// sad error
		//	writeErrorResponse()
		//default:
		//	// suprised error
		//	writeErrorResponse()
		//}
	}
	writeJsonResponse(writer, boxDTO)
}

func (handler *GetBoxRequestHandler) parseBoxKey(url string) string {
	return ""
}

func writeJsonResponse(writer http.ResponseWriter, i interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(i)
}
