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

package commands

import "github.com/urfave/cli"

// Commands holds all previously registered commands.
var Commands []cli.Command

func init() {
	Commands = make([]cli.Command, 0, 4)
}

// Register registers a command at the main app.
func Register(cmd cli.Command) {
	Commands = append(Commands, cmd)
}
