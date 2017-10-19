package main

import (
	"github.com/urfave/cli"
	"os"
)

type CLIParams struct {
	Port int
}

func NewCLIParams() *CLIParams {
	cliParams := &CLIParams{}

	app := cli.NewApp()
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