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

const APP_VER = "0.2.0"

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
