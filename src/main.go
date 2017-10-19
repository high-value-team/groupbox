package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/ralfw/groupbox/src/backend"
)

// wird durch build.sh gesetzt
var VersionNumber string = ""

func main() {
	cliParams := NewCLIParams(VersionNumber)
	mongoDBAdapter := backend.MongoDBAdapter{}
	mongoDBAdapter.Start()
	defer mongoDBAdapter.Stop()

	interactions := backend.NewInteractions(&mongoDBAdapter)
	requestHandlers := []backend.RequestHandler{
		&backend.VersionRequestHandler{VersionNumber: VersionNumber},
		&backend.GetBoxRequestHandler{Interactions: interactions},
		&backend.StaticRequestHandler{},
	}
	httpPortal := backend.HTTPPortal{RequestHandlers: requestHandlers}
	httpPortal.Run(cliParams.Port)
}
