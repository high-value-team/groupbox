package request_handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Version struct {
	VersionNumber string
}

func NewVersionHandler(versionNumber string) http.HandlerFunc {
	version := Version{VersionNumber: versionNumber}
	return version.Handle
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
