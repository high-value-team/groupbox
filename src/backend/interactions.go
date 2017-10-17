package backend

import "time"

type VersionInfo struct {
	VersionNumber string `json:"versionNumber"`
	Timestamp string `json:"timestamp"`
}

func getVersionInformation() *VersionInfo {
	return &VersionInfo{
		VersionNumber: "0.0.1",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
