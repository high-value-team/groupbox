package backend

import (
	"net/http"
	"regexp"
	"encoding/json"
	"github.com/ralfw/groupbox/src/frontend"
)

type RouteType int
const (
	Version RouteType = iota
	// CreateBox
	// OpenBox
	// AddItem
	Frontend
)

type HTTPPortal struct {
}

func (portal *HTTPPortal) Run(address string){
	http.ListenAndServe(address, portal)
}

func (portal *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	routeType := portal.classifyRoute(reader.URL.Path)
	switch routeType {
	case Frontend:
		http.FileServer(frontend.FS(false)).ServeHTTP(writer, reader)
	case Version:
		portal.HandleVersion(writer, reader)
	default:
		portal.HandleNotFound(writer, reader)
	}
}

func (portal *HTTPPortal) classifyRoute(url string) RouteType {
	backendPath := regexp.MustCompile("^/api/([a-zA-Z0-9]+)$")
	backendPathRegex := backendPath.FindStringSubmatch(url)

	if backendPathRegex != nil {
		return Version
	} else {
		return Frontend
	}
}

func (portal *HTTPPortal) HandleVersion(w http.ResponseWriter, r *http.Request) {
	versionInformation := getVersionInformation()
	portal.writeResponse(w, versionInformation)
}

func (portal *HTTPPortal) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (portal *HTTPPortal) writeResponse(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}
