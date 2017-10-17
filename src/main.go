package main

//go:generate go run frontend/util/generator/generator.go

import (
	"fmt"
	"net/http"

	"github.com/ralfw/groupbox/src/backend"
	"github.com/ralfw/groupbox/src/frontend"
)

// usage:
// go run main.go
// visit localhost:8080
func main() {
	frontedHandler := http.FileServer(frontend.FS(false)).ServeHTTP
	backendHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from Backend")
	}
	httpPortal := backend.HTTPPortal{frontedHandler, backendHandler}
	http.ListenAndServe(":8080", &httpPortal)
}
