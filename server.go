package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Run(app *cli.Context) error {
	return run(app)
}

func run(app *cli.Context) error {
	logrus.Info("lzx good")
	return nil
}
