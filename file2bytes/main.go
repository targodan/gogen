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
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/atotto/clipboard"
	"github.com/targodan/gogen/commands"
	"github.com/targodan/gogen/conv"
	"github.com/urfave/cli"
)

type templateParameter struct {
	Output   string
	Filename string
}

func init() {
	commands.Register(cli.Command{
		Name:    "file2bytes",
		Aliases: []string{"f2b"},
		Usage:   "Convert a file into a golang byte slice.",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "out, o",
				Usage: "Output filename.",
			},
			cli.IntFlag{
				Name:  "linebreak, b",
				Usage: "Break line after [value] bytes.",
				Value: 12,
			},
			cli.BoolFlag{
				Name:  "clipboard, c",
				Usage: "Copy output to clipboard.",
			},
			cli.StringFlag{
				Name:  "template, t",
				Usage: "Use a template file where {{.Output}} is replaced with the bytes output and {{.Filename}} is the filename of the input file.",
			},
		},
		Action: run,
	})
}

func run(c *cli.Context) (err error) {
	infile := c.Args().Get(0)
	data, err := conv.FileOrStdin(infile)
	if err != nil {
		return
	}

	out := conv.BytesToString(data, c.Int("linebreak"))

	templateFile := c.String("template")
	if templateFile != "" {
		tmplStr, err := ioutil.ReadFile(templateFile)
		if err != nil {
			return err
		}
		tmpl, err := template.New("Generated").Parse(string(tmplStr))
		if err != nil {
			return err
		}
		var buffer bytes.Buffer
		err = tmpl.Execute(&buffer, templateParameter{out, infile})
		if err != nil {
			return err
		}
		out = buffer.String()
	}

	outname := c.String("out")
	if outname == "" {
		fmt.Println(out)
	} else {
		ioutil.WriteFile(outname, []byte(out), 0644)
	}

	if c.GlobalBool("clipboard") || c.Bool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
