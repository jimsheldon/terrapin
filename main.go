package main

import (
	"fmt"
	"os"

	"github.com/jimsheldon/terrapin/install"

	"github.com/urfave/cli"
)

var version = "undefined"

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "terrapin"
	app.Usage = "install a specific version of Terraform to a desired directory"
	app.Version = version

	app.Commands = []cli.Command{
		install.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
