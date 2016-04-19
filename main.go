package main

import (
	"github.com/codegangsta/cli"
	"github.com/nathan-osman/escapefromlibc/command/net"

	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "elc"
	app.Usage = "system rescue tool"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Nathan Osman",
			Email: "nathan@quickmediasolutions.com",
		},
	}
	app.Commands = []cli.Command{
		net.WgetCommand,
	}
	app.Writer = os.Stderr
	app.Run(os.Args)
}
