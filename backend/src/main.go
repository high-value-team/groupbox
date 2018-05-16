package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/high-value-team/groupbox/backend/src/interior_interactions"
	"github.com/high-value-team/groupbox/backend/src/portal_http"
	"github.com/high-value-team/groupbox/backend/src/provider_mongodb"
	"github.com/high-value-team/groupbox/backend/src/provider_smtp"
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
	httpPortal := portal_http.NewHTTPPortal(interactions, VersionNumber)

	// run
	httpPortal.Run(cliParams.Port)
}

func NewProviders(cliParams *CLIParams) (*provider_mongodb.MongoDBAdapter, *provider_smtp.EmailNotifications) {
	mongoDBAdapter := provider_mongodb.MongoDBAdapter{ConnectionString: cliParams.MongoDBURL}
	mongoDBAdapter.Start()

	emailNotifications := &provider_smtp.EmailNotifications{
		Domain:            cliParams.Domain,
		NoReplyEmail:      cliParams.SMTPNoReplyEmail,
		SMTPServerAddress: cliParams.SMTPServerAddress,
		Username:          cliParams.SMTPUsername,
		Password:          cliParams.SMTPPassword,
	}
	return &mongoDBAdapter, emailNotifications
}

func NewInteractions(mongoDBAdapter *provider_mongodb.MongoDBAdapter, emailNotifications *provider_smtp.EmailNotifications) *interior_interactions.Interactions {
	return interior_interactions.NewInteractions(mongoDBAdapter, emailNotifications)
}
