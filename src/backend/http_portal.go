package backend

import (
	"net/http"
	"regexp"
)

var backendPath = regexp.MustCompile("^/api/([a-zA-Z0-9]+)$")

type HTTPPortal struct {
	FrontendHandler http.HandlerFunc
	BackendHandler  http.HandlerFunc
}

func (p *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	onFrontendRoute := func() {
		p.FrontendHandler(writer, reader)
	}
	onBackendRoute := func() {
		p.BackendHandler(writer, reader)
	}
	p.route(reader, onFrontendRoute, onBackendRoute)
}

func (p *HTTPPortal) route(reader *http.Request, onFrontendRoute, onBackendRoute func()) {
	backendPathRegex := backendPath.FindStringSubmatch(reader.URL.Path)

	if backendPathRegex != nil {
		onBackendRoute()
	} else {
		onFrontendRoute()
	}
}
