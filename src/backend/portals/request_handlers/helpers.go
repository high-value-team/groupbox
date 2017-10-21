package request_handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/high-value-team/groupbox/src/backend/models"
)

func match(regexStr string, path string) bool {
	regex := regexp.MustCompile(regexStr)
	if regex.FindStringSubmatch(path) != nil {
		return true
	}
	return false
}

func writeJsonResponse(writer http.ResponseWriter, i interface{}) {
	writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(i)
}

func parseRequestBody(reader *http.Request, body interface{}) {
	decoder := json.NewDecoder(reader.Body)
	defer func() {
		err := reader.Body.Close()
		if err != nil {
			panic(models.SuprisingException{Err: err})
		}
	}()
	err := decoder.Decode(&body)
	if err != nil {
		panic(models.SuprisingException{Err: err})
	}
}
