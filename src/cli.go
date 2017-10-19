package main

import (
	"os"

	"github.com/urfave/cli"
)

type CLIParams struct {
	Port int
}

func NewCLIParams(version string) *CLIParams {
	cliParams := &CLIParams{}

	app := cli.NewApp()
	app.Version = version
	const FlagHTTPPort string = "port"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   FlagHTTPPort,
			EnvVar: "HTTP_PORT",
			Value:  8080,
			Usage:  "http port",
		},
	}
	app.Action = func(c *cli.Context) error {
		cliParams.Port = c.Int(FlagHTTPPort)
		return nil
	}
	app.Run(os.Args)

	return cliParams
}
