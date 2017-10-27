package request_handlers

import (
	"net/http"

	"github.com/high-value-team/groupbox/src/frontend"
)

type StaticContent struct{}

func NewStaticContentHandler() http.HandlerFunc {
	staticContent := StaticContent{}
	return staticContent.Handle
}

func (handler *StaticContent) Handle(writer http.ResponseWriter, reader *http.Request) {
	if !frontend.IsExist(reader.URL.Path) {
		reader.URL.Path = "/"
	}
	http.FileServer(frontend.FS(false)).ServeHTTP(writer, reader)
}
