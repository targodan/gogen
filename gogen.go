package main

import (
	"os"
	"runtime"

	"github.com/targodan/gogen/file2bytes"
	"github.com/urfave/cli"
)

const APP_VER = "0.1.0-dev"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "gogen"
	app.Usage = "A nifty tool to generate snippets for go."
	app.Version = APP_VER

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "clipboard, c",
			Usage: "Copy output to clipboard",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "file2bytes",
			Aliases: []string{"f2b"},
			Usage:   "Convert a file into a golang byte slice.",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "linebreak, b",
					Usage: "break line after n bytes",
					Value: 12,
				},
			},
			Action: file2bytes.Run,
		},
	}

	app.Run(os.Args)
}
