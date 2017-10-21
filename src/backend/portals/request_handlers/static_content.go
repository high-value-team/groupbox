package request_handlers

import (
	"net/http"

	"github.com/high-value-team/groupbox/src/frontend"
)

type StaticContent struct{}

func (handler *StaticContent) TryHandle(writer http.ResponseWriter, reader *http.Request) bool {
	if !frontend.IsExist(reader.URL.Path) {
		reader.URL.Path = "/"
	}
	http.FileServer(frontend.FS(false)).ServeHTTP(writer, reader)
	return true
}
