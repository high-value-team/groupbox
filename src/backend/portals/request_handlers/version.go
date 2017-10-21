package request_handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Version struct {
	VersionNumber string
}

func (handler *Version) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	if handler.Match(reader) {
		handler.Handle(writer, reader)
		return true
	}
	return false
}

func (handler *Version) Match(reader *http.Request) bool {
	return match("^/api/version$", reader.URL.Path)
}

func (handler *Version) Handle(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
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