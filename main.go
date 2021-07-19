package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := NewApp()
	app.Commands = []cli.Command{
		NewServerCommand(Run),
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
