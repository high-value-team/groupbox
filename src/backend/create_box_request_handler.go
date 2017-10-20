package backend

import (
	"encoding/json"
	"net/http"
)

type CreateBoxRequestHandler struct {
	Interactions *Interactions
}

type CreateBoxResponseDTO struct {
	BoxKey string `json:"boxKey"`
}
type CreateBoxRequestDTO struct {
	Title   string   `json:"title"`
	Owner   string   `json:"owner"`
	Members []string `json:"members"`
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
	requestDTO := CreateBoxRequestDTO{}
	parseRequestBody(reader, &requestDTO)
	responseDTO := handler.Interactions.CreateBox(requestDTO.Title, requestDTO.Owner, requestDTO.Members)
	writeJsonResponse(writer, responseDTO)
}

func parseRequestBody(reader *http.Request, body interface{}) {
	decoder := json.NewDecoder(reader.Body)
	defer func() {
		err := reader.Body.Close()
		if err != nil {
			panic(SuprisingException{Err: err})
		}
	}()
	err := decoder.Decode(&body)
	if err != nil {
		panic(SuprisingException{Err: err})
	}
}
