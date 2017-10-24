package main

//go:generate go run frontend/util/generator/generator.go

import (
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
	httpPortal := NewHTTPPortal(interactions)

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

func NewHTTPPortal(interactions *interactions.Interactions) *portals.HTTPPortal {
	requestHandlers := []portals.RequestHandler{
		&request_handlers.AddItem{Interactions: interactions},
		&request_handlers.CreateBox{Interactions: interactions},
		&request_handlers.GetBox{Interactions: interactions},
		&request_handlers.Version{VersionNumber: VersionNumber},
		&request_handlers.StaticContent{},
	}
	return &portals.HTTPPortal{RequestHandlers: requestHandlers}
}

func NewInteractions(mongoDBAdapter *providers.MongoDBAdapter, emailNotifications *providers.EmailNotifications) *interactions.Interactions {
	return interactions.NewInteractions(mongoDBAdapter, emailNotifications)
}
