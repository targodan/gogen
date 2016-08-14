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
