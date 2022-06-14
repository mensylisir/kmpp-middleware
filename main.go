package main

import (
	"fmt"
	"github.com/mensylisir/kmpp-middleware/src/server"
	"github.com/toolkits/pkg/runner"
	"github.com/urfave/cli/v2"
	"os"
)

var VERSION = "not specified"

func printEnv() {
	runner.Init()
	fmt.Println("runner.cwd:", runner.Cwd)
	fmt.Println("runner.hostname:", runner.Hostname)
	fmt.Println("runner.fd_limits:", runner.FdLimits())
	fmt.Println("runner.vm_limits:", runner.VMLimits())
}

func NewServerCmd() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run Server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "conf",
				Aliases: []string{"c"},
				Usage:   "specify configuration file(.json,.yaml,.toml)",
			},
		},
		Action: func(ctx *cli.Context) error {
			printEnv()
			var options []server.ServerOption
			if ctx.String("conf") != "" {
				options = append(options, server.SetConfigFile(ctx.String("conf")))
			}
			options = append(options, server.SetVersion(VERSION))
			server.Run(options...)
			return nil
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "MiddleWare"
	app.Version = VERSION
	app.Usage = "MiddleWare"
	app.Commands = []*cli.Command{
		NewServerCmd(),
	}
	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
