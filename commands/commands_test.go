package commands

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/cli"
)

func TestRegister(t *testing.T) {
	Convey("Registering commands should work.", t, func() {
		cmd := cli.Command{
			Name:    "dummyfortesting",
			Aliases: []string{"dft"},
			Usage:   "Just testing.",
		}

		Register(cmd)
		So(Commands, ShouldContain, cmd)
	})
}
