package base64

import (
	"encoding/base64"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/targodan/gogen/commands"
	"github.com/targodan/gogen/conv"
	"github.com/urfave/cli"
)

func init() {
	commands.Register(cli.Command{
		Name:  "base64",
		Usage: "Encode and decode base64 data.",
		Subcommands: []cli.Command{
			{
				Name:  "decode",
				Usage: "Decode a base64 string.",
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
				Action: decode,
			},
			{
				Name:  "encode",
				Usage: "Encode bytes as a base64 string.",
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
				Action: encode,
			},
		},
	})
}

func decode(c *cli.Context) (err error) {
	text, err := conv.TextOrStdin(c.Args().Get(0))
	if err != nil {
		return
	}

	data, err := base64.StdEncoding.DecodeString(text)
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

func encode(c *cli.Context) (err error) {
	text, err := conv.TextOrStdin(c.Args().Get(0))
	if err != nil {
		return
	}

	data, err := conv.TextToByteSlice(text)
	if err != nil {
		return
	}

	out := base64.StdEncoding.EncodeToString(data)
	fmt.Println(out)

	if c.GlobalBool("clipboard") || c.Bool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
