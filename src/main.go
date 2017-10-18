package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/ralfw/groupbox/src/backend"
)

// wird durch build.sh gesetzt
var VersionNumber string = ""

func main() {
	cliParams := NewCLIParams()
	requestHandlers := []backend.RequestHandler{
		&backend.VersionRequestHandler{VersionNumber: VersionNumber},
		&backend.StaticRequestHandler{},
	}
	httpPortal := backend.HTTPPortal{RequestHandlers: requestHandlers}
	httpPortal.Run(cliParams.Port)
}
