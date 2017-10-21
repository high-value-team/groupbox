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
	// choose one
	//main1()
	//main2()
	main3()
}

// - SLA
func main1() {
	cliParams := NewCLIParams(VersionNumber)

	// MongoDB
	mongoDBAdapter := providers.MongoDBAdapter{ConnectionString: cliParams.MongoDBURL}
	mongoDBAdapter.Start()
	defer mongoDBAdapter.Stop()

	// Email
	emailNotifications := providers.EmailNotifications{
		Domain:            cliParams.Domain,
		NoReplyEmail:      cliParams.SMTPNoReplyEmail,
		SMTPServerAddress: cliParams.SMTPServerAddress,
		Username:          cliParams.SMTPUsername,
		Password:          cliParams.SMTPPassword,
	}

	// Request Handlers
	interactions := interactions.NewInteractions(&mongoDBAdapter, &emailNotifications)
	requestHandlers := []portals.RequestHandler{
		&request_handlers.AddItem{Interactions: interactions},
		&request_handlers.CreateBox{Interactions: interactions},
		&request_handlers.GetBox{Interactions: interactions},
		&request_handlers.Version{VersionNumber: VersionNumber},
		&request_handlers.StaticContent{},
	}

	// HTTP-Portal
	httpPortal := portals.HTTPPortal{RequestHandlers: requestHandlers}
	httpPortal.Run(cliParams.Port)
}

// - zu viele Zeilen
func main2() {
	cliParams := NewCLIParams(VersionNumber)

	mongoDBAdapter := NewMongoDBAdapter(cliParams)
	emailNotifications := NewEmailNotifications(cliParams)

	interactions := NewInteractions(mongoDBAdapter, emailNotifications)

	requestHandlers := NewRequestHandlers(interactions)
	httpPortal := NewHTTPPortal(requestHandlers)

	httpPortal.Run(cliParams.Port)
}

func main3() {
	// config
	cliParams := NewCLIParams(VersionNumber)

	// consturct
	mongoDBAdapter, emailNotifications := NewProviders(cliParams)
	interactions := NewInteractions(mongoDBAdapter, emailNotifications)
	httpPortal := NewHTTPPortal2(interactions)

	// run
	httpPortal.Run(cliParams.Port)
}

func NewProviders(cliParams *CLIParams) (*providers.MongoDBAdapter, *providers.EmailNotifications) {
	mongoDBAdapter := providers.MongoDBAdapter{ConnectionString: cliParams.MongoDBURL}
	mongoDBAdapter.Start()
	defer mongoDBAdapter.Stop()

	emailNotifications := &providers.EmailNotifications{
		Domain:            cliParams.Domain,
		NoReplyEmail:      cliParams.SMTPNoReplyEmail,
		SMTPServerAddress: cliParams.SMTPServerAddress,
		Username:          cliParams.SMTPUsername,
		Password:          cliParams.SMTPPassword,
	}
	return &mongoDBAdapter, emailNotifications
}

func NewHTTPPortal2(interactions *interactions.Interactions) *portals.HTTPPortal {
	requestHandlers := []portals.RequestHandler{
		&request_handlers.AddItem{Interactions: interactions},
		&request_handlers.CreateBox{Interactions: interactions},
		&request_handlers.GetBox{Interactions: interactions},
		&request_handlers.Version{VersionNumber: VersionNumber},
		&request_handlers.StaticContent{},
	}
	return &portals.HTTPPortal{RequestHandlers: requestHandlers}
}

func NewMongoDBAdapter(cliParams *CLIParams) *providers.MongoDBAdapter {
	mongoDBAdapter := providers.MongoDBAdapter{ConnectionString: cliParams.MongoDBURL}
	mongoDBAdapter.Start()
	defer mongoDBAdapter.Stop()
	return &mongoDBAdapter
}

func NewEmailNotifications(cliParams *CLIParams) *providers.EmailNotifications {
	return &providers.EmailNotifications{
		Domain:            cliParams.Domain,
		NoReplyEmail:      cliParams.SMTPNoReplyEmail,
		SMTPServerAddress: cliParams.SMTPServerAddress,
		Username:          cliParams.SMTPUsername,
		Password:          cliParams.SMTPPassword,
	}
}

func NewInteractions(mongoDBAdapter *providers.MongoDBAdapter, emailNotifications *providers.EmailNotifications) *interactions.Interactions {
	return interactions.NewInteractions(mongoDBAdapter, emailNotifications)
}

func NewRequestHandlers(interactions *interactions.Interactions) []portals.RequestHandler {
	return []portals.RequestHandler{
		&request_handlers.AddItem{Interactions: interactions},
		&request_handlers.CreateBox{Interactions: interactions},
		&request_handlers.GetBox{Interactions: interactions},
		&request_handlers.Version{VersionNumber: VersionNumber},
		&request_handlers.StaticContent{},
	}
}

func NewHTTPPortal(requestHandlers []portals.RequestHandler) *portals.HTTPPortal {
	return &portals.HTTPPortal{RequestHandlers: requestHandlers}
}
