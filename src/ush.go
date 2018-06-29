package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"log"

	"github.com/mkideal/cli"
)

var (
	appVersion = "0.0.1"
	appLicense = "MIT License"

	appLicenseDetail = "A short and simple permissive license with conditions only requiring \npreservation of copyright and license notices. Licensed works, \nmodifications, and larger works may be distributed under different \nterms and without source code."
	appCopyrightYear = 2018
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(pwd),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("Provides Help information for built-in commands")

type rootT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"show version number and quit"`
	License bool `cli:"l,license" usage:"show license and quit"`
}

var root = &cli.Command{
	Desc: "Welcome to UniversalShell (ush)",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			currentYear := time.Now().UTC().Year()
			copyrightDuration := strconv.Itoa(appCopyrightYear)
			if currentYear > appCopyrightYear {
				copyrightDuration += "-" + strconv.Itoa(currentYear)
			}
			ctx.String("UniversalShell (ush)    version %s\nCopyright (c) %s Abhishek Kumar.\nIt is licensed under the '%s'.\n\n", appVersion, copyrightDuration, appLicense)
		}
		if argv.License {
			ctx.String("The 'UniversalShell (ush)' is licensed under the '%s'\n\n%s\n\n", appLicense, appLicenseDetail)
		}
		return nil
	},
}

type pwdT struct {
	cli.Helper
}

var pwd = &cli.Command{
	Name: "pwd",
	Desc: "Displays the path of present working directory",
	Argv: func() interface{} { return new(pwdT) },
	Fn: func(ctx *cli.Context) error {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		ctx.String("%s\n", dir)
		return nil
	},
}