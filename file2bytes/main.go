package file2bytes

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/targodan/gogen/commands"
	"github.com/targodan/gogen/conv"
	"github.com/urfave/cli"
)

func init() {
	commands.Register(cli.Command{
		Name:    "file2bytes",
		Aliases: []string{"f2b"},
		Usage:   "Convert a file into a golang byte slice.",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "linebreak, b",
				Usage: "Break line after [value] bytes.",
				Value: 12,
			},
			cli.BoolFlag{
				Name:  "clipboard, c",
				Usage: "Copy output to clipboard.",
			},
		},
		Action: run,
	})
}

func run(c *cli.Context) (err error) {
	data, err := conv.FileOrStdin(c.Args().Get(0))
	if err != nil {
		return
	}

	out := conv.BytesToString(data, c.Int("linebreak"))
	fmt.Println(out)

	if c.GlobalBool("clipboard") || c.Bool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
