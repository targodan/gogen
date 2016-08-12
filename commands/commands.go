package commands

import "github.com/urfave/cli"

var Commands []cli.Command

func init() {
	Commands = make([]cli.Command, 0, 4)
}

func Register(cmd cli.Command) {
	Commands = append(Commands, cmd)
}
