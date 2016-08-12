package file2bytes

import (
	"fmt"
	"io/ioutil"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "file2bytes",
		Aliases: []string{"f2b"},
		Usage:   "Convert a file into a golang byte slice.",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "linebreak, b",
				Usage: "break line after n bytes",
				Value: 12,
			},
			cli.BoolFlag{
				Name:  "clipboard, c",
				Usage: "Copy output to clipboard",
			},
		},
		Action: run,
	}
}

func run(c *cli.Context) (err error) {
	filename := c.Args().Get(0)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	out := "[]byte{"
	ln := 0
	for _, b := range data {
		out = fmt.Sprintf("%s0x%02x, ", out, b)
		ln++
		if ln >= c.Int("linebreak") {
			out = fmt.Sprintln(out)
			ln = 0
		}
	}
	out = out[:len(out)-2] + "}"
	fmt.Println(out)

	if c.GlobalBool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
