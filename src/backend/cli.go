package backend

import "github.com/urfave/cli"

const (
	FlagHTTPPort string = "port"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:   FlagHTTPPort,
			EnvVar: "HTTP_PORT",
			Value:  8080,
			Usage:  "http port",
		},
	}
}
