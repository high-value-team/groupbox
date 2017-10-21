package portals

import (
	"fmt"
	"net/http"
	"time"
	"github.com/high-value-team/groupbox/src/backend/models"
)

type RequestHandler interface {
	TryHandle(writer http.ResponseWriter, reader *http.Request) bool
}

type HTTPPortal struct {
	RequestHandlers []RequestHandler
}

func (portal *HTTPPortal) Run(port int) {
	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, portal)
}

func (portal *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Cache-Control", "no-cache")
	defer handleException(writer)
	for _, requestHandler := range portal.RequestHandlers {
		if requestHandler.TryHandle(writer, reader) {
			break
		}
	}
}

func handleException(writer http.ResponseWriter) {
	writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	if r := recover(); r != nil {
		switch ex := r.(type) {
		case models.SadException:
			http.Error(writer, ex.Message(), 404)
		case models.SuprisingException:
			http.Error(writer, ex.Message(), 500)
		default:
			if err, ok := r.(error); ok {
				http.Error(writer, err.Error(), 500)
			} else {
				http.Error(writer, fmt.Sprintf("%s", r), 500)
			}
		}
	}
}
