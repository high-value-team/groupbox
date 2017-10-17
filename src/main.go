package main

import (
	"github.com/ralfw/groupbox/src/backend"
	"net/http"
	"fmt"
)

// usage:
// go run main.go
// visit localhost:8080
func main() {
	frontedHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"hello from Frontend")
	}
	backendHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,"hello from Backend")
	}
	httpPortal := portal.HTTPPortal{frontedHandler, backendHandler}
	http.ListenAndServe(":8080", &httpPortal)
}
