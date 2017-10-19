package main

import (
	"os"

	"github.com/urfave/cli"
)

type CLIParams struct {
	Port       int
	MongoDBURL string
}

func NewCLIParams() *CLIParams {
	cliParams := &CLIParams{}

	app := cli.NewApp()
	const FlagHTTPPort string = "port"
	const FlagMongoDBURL string = "mongodb-url"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   FlagHTTPPort,
			EnvVar: "HTTP_PORT",
			Value:  8080,
			Usage:  "http port",
		},
		cli.StringFlag{
			Name:   FlagMongoDBURL,
			EnvVar: "MONGODB_URL",
			Value:  "mongodb://localhost:27017/develop",
			Usage:  "MongoDB URL",
		},
	}
	app.Action = func(c *cli.Context) error {
		cliParams.Port = c.Int(FlagHTTPPort)
		cliParams.MongoDBURL = c.String(FlagMongoDBURL)
		return nil
	}
	app.Run(os.Args)

	return cliParams
}
