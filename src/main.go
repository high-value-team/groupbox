package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/ralfw/groupbox/src/backend"
)

// wird durch build.sh gesetzt
var VersionNumber string = ""

func main() {
	cliParams := NewCLIParams()
	interactions := &backend.Interactions{VersionNumber: VersionNumber}
	httpPortal := backend.HTTPPortal{Interactions: interactions}
	httpPortal.Run(cliParams.Port)
}
