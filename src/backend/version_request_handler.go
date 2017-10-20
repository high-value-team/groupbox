package backend

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"
)

type VersionRequestHandler struct {
	VersionNumber string
}

func (handler *VersionRequestHandler) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	if handler.Match(reader) {
		handler.Handle(writer, reader)
		return true
	}
	return false
}

func (handler *VersionRequestHandler) Match(reader *http.Request) bool {
	versionPath := regexp.MustCompile("^/api/([a-zA-Z0-9]+)$")
	return versionPath.FindStringSubmatch(reader.URL.Path) != nil
}

func (handler *VersionRequestHandler) Handle(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(VersionInfo{
		VersionNumber: handler.VersionNumber,
		Timestamp:     time.Now().Format(time.RFC3339),
	})
}

type VersionInfo struct {
	VersionNumber string `json:"versionNumber"`
	Timestamp     string `json:"timestamp"`
}
