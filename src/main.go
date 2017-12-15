package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/go-chi/chi"
	"github.com/high-value-team/groupbox/src/backend/interior/interactions"
	"github.com/high-value-team/groupbox/src/backend/portals"
	"github.com/high-value-team/groupbox/src/backend/portals/request_handlers"
	"github.com/high-value-team/groupbox/src/backend/providers"
)

// wird durch build.sh gesetzt
var VersionNumber string = ""

func main() {
	// config
	cliParams := NewCLIParams(VersionNumber)

	// consturct
	mongoDBAdapter, emailNotifications := NewProviders(cliParams)
	defer mongoDBAdapter.Stop()
	interactions := NewInteractions(mongoDBAdapter, emailNotifications)
	httpPortal := NewHTTPPortal(interactions, VersionNumber)

	// run
	httpPortal.Run(cliParams.Port)
}

func NewProviders(cliParams *CLIParams) (*providers.MongoDBAdapter, *providers.EmailNotifications) {
	mongoDBAdapter := providers.MongoDBAdapter{ConnectionString: cliParams.MongoDBURL}
	mongoDBAdapter.Start()

	emailNotifications := &providers.EmailNotifications{
		Domain:            cliParams.Domain,
		NoReplyEmail:      cliParams.SMTPNoReplyEmail,
		SMTPServerAddress: cliParams.SMTPServerAddress,
		Username:          cliParams.SMTPUsername,
		Password:          cliParams.SMTPPassword,
	}
	return &mongoDBAdapter, emailNotifications
}

func NewHTTPPortal(interactions *interactions.Interactions, versionNumber string) *portals.HTTPPortal {
	return &portals.HTTPPortal{Router: NewRouter(interactions, versionNumber)}
}

func NewInteractions(mongoDBAdapter *providers.MongoDBAdapter, emailNotifications *providers.EmailNotifications) *interactions.Interactions {
	return interactions.NewInteractions(mongoDBAdapter, emailNotifications)
}

func NewRouter(interactions *interactions.Interactions, versionNumber string) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/api/boxes/{boxKey}/items", request_handlers.NewAddItemHandler(interactions))
	router.Put("/api/boxes/{boxKey}/items/{itemID}", request_handlers.NewUpdateItemHandler(interactions))
	router.Delete("/api/boxes/{boxKey}/items/{itemID}", request_handlers.NewDeleteItemHandler(interactions))
	router.Post("/api/boxes", request_handlers.NewCreateBoxHandler(interactions))
	router.Get("/api/boxes/{boxKey}", request_handlers.NewGetBoxHandler(interactions))
	router.Get("/api/version", request_handlers.NewVersionHandler(versionNumber))
	router.NotFound(request_handlers.NewStaticContentHandler())
	return router
}
