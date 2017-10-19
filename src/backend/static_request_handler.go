package backend

import (
	"net/http"

	"github.com/high-value-team/groupbox/src/frontend"
)

type StaticRequestHandler struct{}

func (handler *StaticRequestHandler) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	http.FileServer(frontend.FS(false)).ServeHTTP(writer, reader)
	return true
}
