package base64

import (
	"encoding/hex"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/targodan/gogen/commands"
	"github.com/urfave/cli"
)

func init() {
	commands.Register(cli.Command{
		Name:  "hex",
		Usage: "encode and decode hex",
		Subcommands: []cli.Command{
			{
				Name:  "decode",
				Usage: "decode a hex string",
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
				Action: decode,
			},
		},
	})
}

func decode(c *cli.Context) (err error) {
	text := c.Args().Get(0)

	data, err := hex.DecodeString(text)
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
	if ln == 0 {
		out = out[:len(out)-1]
	}
	out = out[:len(out)-2] + "}"
	fmt.Println(out)

	if c.GlobalBool("clipboard") || c.Bool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
