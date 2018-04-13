package main

//go:generate go run frontend/util/generator/generator.go

import (
	"github.com/high-value-team/groupbox/backend/src/interior/interactions"
	"github.com/high-value-team/groupbox/backend/src/portals"
	"github.com/high-value-team/groupbox/backend/src/providers"
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
	httpPortal := portals.NewHTTPPortal(interactions, VersionNumber)

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

func NewInteractions(mongoDBAdapter *providers.MongoDBAdapter, emailNotifications *providers.EmailNotifications) *interactions.Interactions {
	return interactions.NewInteractions(mongoDBAdapter, emailNotifications)
}
