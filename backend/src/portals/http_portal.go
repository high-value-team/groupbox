package portals

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/backend/src/exceptions"
	"github.com/high-value-team/groupbox/backend/src/interior/interactions"
	"github.com/high-value-team/groupbox/backend/src/portals/request_handlers"
	"github.com/rs/cors"
)

type HTTPPortal struct {
	router *chi.Mux
}

func NewHTTPPortal(interactions *interactions.Interactions, versionNumber string) *HTTPPortal {

	router := chi.NewRouter()
	router.Post("/api/boxes/{boxKey}/items", request_handlers.NewAddItemHandler(interactions))
	router.Put("/api/boxes/{boxKey}/items/{itemID}", request_handlers.NewUpdateItemHandler(interactions))
	router.Delete("/api/boxes/{boxKey}/items/{itemID}", request_handlers.NewDeleteItemHandler(interactions))
	router.Post("/api/boxes", request_handlers.NewCreateBoxHandler(interactions))
	router.Get("/api/boxes/{boxKey}", request_handlers.NewGetBoxHandler(interactions))
	router.Get("/api/version", request_handlers.NewVersionHandler(versionNumber))
	//router.NotFound(request_handlers.NewStaticContentHandler())

	return &HTTPPortal{router: router}
}

func (portal *HTTPPortal) Run(port int) {
	address := fmt.Sprintf(":%d", port)
	handler := cors.AllowAll().Handler(portal)

	log.Println("Starting Webserver, please go to: ", address)
	http.ListenAndServe(address, handler)
}

func (portal *HTTPPortal) ServeHTTP(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	defer handleException(writer)
	portal.router.ServeHTTP(writer, reader)
}

func handleException(writer http.ResponseWriter) {
	r := recover()
	if r != nil {
		switch ex := r.(type) {
		case exceptions.SadException:
			http.Error(writer, ex.Message(), 404)
		case exceptions.SuprisingException:
			http.Error(writer, ex.Message(), 500)
		default:
			if err, ok := r.(error); !ok {
				http.Error(writer, err.Error(), 500)
			} else {
				http.Error(writer, fmt.Sprintf("%s", r), 500)
			}
		}
	}
}
