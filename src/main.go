package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/ralfw/groupbox/src/backend"
)

// usage:
// go run main.go
func main() {
	httpPortal := backend.HTTPPortal{}
	httpPortal.Run(":80")
}
