/* Copyright (C) 2016 Luca Corbatto
 *
 * This file is part of the gogen project.
 *
 * The gogen project is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The gogen project is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package base64

import (
	"encoding/hex"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/targodan/gogen/commands"
	"github.com/targodan/gogen/conv"
	"github.com/urfave/cli"
)

func init() {
	commands.Register(cli.Command{
		Name:  "hex",
		Usage: "Encode and decode hex data.",
		Subcommands: []cli.Command{
			{
				Name:  "decode",
				Usage: "Decode a hex string.",
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
						Value: 64,
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

	data, err := hex.DecodeString(text)
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

	tmp := hex.EncodeToString(data)
	lnbr := c.Int("linebreak")
	out := ""
	for len(tmp) > 0 {
		partLen := lnbr
		if len(tmp) < lnbr {
			partLen = len(tmp)
		}
		out += fmt.Sprintln(tmp[:partLen])
		tmp = tmp[partLen:]
	}
	fmt.Println(out)

	if c.GlobalBool("clipboard") || c.Bool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
