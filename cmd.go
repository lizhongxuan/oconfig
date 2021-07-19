package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

var (
	appName   = filepath.Base(os.Args[0])
	Debug     bool
	DebugFlag = cli.BoolFlag{
		Name:        "debug",
		Usage:       "(logging) Turn on debug logs",
		Destination: &Debug,
	}
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "config center"
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
	}
	app.Flags = []cli.Flag{
		DebugFlag,
	}
	app.Before = SetupDebug(nil)

	return app
}

func NewCommand(action func(*cli.Context) error) cli.Command {
	return cli.Command{
		Name:            "ctr",
		Usage:           "Run ctr",
		SkipFlagParsing: true,
		SkipArgReorder:  true,
		Action:          action,
	}
}

func SetupDebug(next func(ctx *cli.Context) error) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		if next != nil {
			return next(ctx)
		}
		return nil
	}
}

const ServerCommand = "server"

var OConfigName = "lzx"

func NewServerCommand(action func(*cli.Context) error) cli.Command {
	return cli.Command{
		Name:            ServerCommand,
		Usage:           "Trigger  a server of config center.",
		SkipFlagParsing: false,
		SkipArgReorder:  true,
		Action:          action,
	}
}
