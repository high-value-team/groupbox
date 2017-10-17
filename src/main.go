package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/ralfw/groupbox/src/backend"
	"github.com/urfave/cli"
	"os"
)

var VersionNumber string = ""

func main() {
	// TODO CLI-Parsing überarbeiten
	app := cli.NewApp()
	app.Flags = backend.Flags()
	app.Action = func( c *cli.Context) error {
		interactions := &backend.Interactions{VersionNumber:VersionNumber}
		httpPortal := backend.HTTPPortal{Interactions:interactions}
		httpPortal.Run(c.Int(backend.FlagHTTPPort))
		return nil
	}
	app.Run(os.Args)
}