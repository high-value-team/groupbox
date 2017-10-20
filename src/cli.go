package main

import (
	"os"

	"github.com/urfave/cli"
)

type CLIParams struct {
	Port              int
	Domain            string
	MongoDBURL        string
	SMTPNoReplyEmail  string
	SMTPServerAddress string
	SMTPUsername      string
	SMTPPassword      string
}

func NewCLIParams(version string) *CLIParams {
	cliParams := &CLIParams{}

	app := cli.NewApp()
	app.Version = version
	const FlagHTTPPort string = "port"
	const FlagDomain string = "groupbox-root-uri"
	const FlagMongoDBURL string = "mongodb-url"
	const FlagSMTPNoReplyEmail string = "smtp-no-reply-email"
	const FlagSMTPServerAddress string = "smtp-server-address"
	const FlagSMTPUsername string = "smtp-username"
	const FlagSMTPPassword string = "smtp-password"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   FlagHTTPPort,
			EnvVar: "HTTP_PORT",
			Value:  8080,
			Usage:  "http port",
		},
		cli.StringFlag{
			Name:   FlagDomain,
			EnvVar: "GROUPBOX_ROOT_URI",
			Value:  "http://localhost:8080",
			Usage:  "Domain of the application",
		},
		cli.StringFlag{
			Name:   FlagMongoDBURL,
			EnvVar: "MONGODB_URL",
			Value:  "mongodb://localhost:27017/develop",
			Usage:  "MongoDB URL",
		},
		cli.StringFlag{
			Name:   FlagSMTPNoReplyEmail,
			EnvVar: "SMTP_NO_REPLY_EMAIL",
			Value:  "no-reply@googlemail.com",
			Usage:  "no-reply email address",
		},
		cli.StringFlag{
			Name:   FlagSMTPServerAddress,
			EnvVar: "SMTP_SERVER_ADDRESS",
			Value:  "localhost:587",
			Usage:  "SMTP server address",
		},
		cli.StringFlag{
			Name:   FlagSMTPUsername,
			EnvVar: "SMTP_USERNAME",
			Value:  "smtp_username",
			Usage:  "SMPT Username",
		},
		cli.StringFlag{
			Name:   FlagSMTPPassword,
			EnvVar: "SMTP_PASSWORD",
			Value:  "smtp_password",
			Usage:  "SMTP Password",
		},
	}
	app.Action = func(c *cli.Context) error {
		cliParams.Port = c.Int(FlagHTTPPort)
		cliParams.Domain = c.String(FlagDomain)
		cliParams.MongoDBURL = c.String(FlagMongoDBURL)
		cliParams.SMTPNoReplyEmail = c.String(FlagSMTPNoReplyEmail)
		cliParams.SMTPServerAddress = c.String(FlagSMTPServerAddress)
		cliParams.SMTPUsername = c.String(FlagSMTPUsername)
		cliParams.SMTPPassword = c.String(FlagSMTPPassword)
		return nil
	}
	app.Run(os.Args)

	return cliParams
}
