package main

import (
	"os"
	"runtime"

	"github.com/targodan/gogen/file2bytes"
	"github.com/urfave/cli"
)

const APP_VER = "0.1.0-dev"

var commands []cli.Command

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "gogen"
	app.Usage = "A nifty tool to generate snippets for go."
	app.Version = APP_VER

	app.Commands = []cli.Command{
		file2bytes.Command(),
	}

	app.Run(os.Args)
}
