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
