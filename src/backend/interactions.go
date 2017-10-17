package backend

import "time"

type VersionInfo struct {
	VersionNumber string `json:"versionNumber"`
	Timestamp string `json:"timestamp"`
}

type Interactions struct {
	VersionNumber string
}

func (i *Interactions) getVersionInformation() *VersionInfo {
	return &VersionInfo{
		VersionNumber: i.VersionNumber,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
