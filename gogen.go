package main

import (
	"os"
	"runtime"

	_ "github.com/targodan/gogen/base64"
	"github.com/targodan/gogen/commands"
	_ "github.com/targodan/gogen/file2bytes"
	_ "github.com/targodan/gogen/hex"

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

	app.Commands = commands.Commands

	app.Run(os.Args)
}
