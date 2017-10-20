package backend

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"
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
