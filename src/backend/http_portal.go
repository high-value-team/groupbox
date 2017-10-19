package backend

import (
	"net/http"
	"fmt"
)

type RequestHandler interface {
	TryHandle(writer http.ResponseWriter, reader *http.Request) bool
}

type HTTPPortal struct {
	RequestHandlers []RequestHandler
}

func (portal *HTTPPortal) Run(port int){
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, portal)
}

func (portal *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	for _, requestHandler := range portal.RequestHandlers {
		if requestHandler.TryHandle(writer, reader) {
			break
		}
	}
}