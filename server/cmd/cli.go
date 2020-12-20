package cmd

import (
	"os"

	"github.com/timvosch/hobby-rl/cmd/serve"
	"github.com/urfave/cli"
)

// Execute the cli
func Execute() {
	app := &cli.App{
		Name: "RLServer",
		Commands: []cli.Command{
			serve.CreateCommand(),
		},
	}

	app.Run(os.Args)
}
