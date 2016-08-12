package file2bytes

import (
	"fmt"
	"io/ioutil"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) (err error) {
	filename := c.Args().Get(0)

	data, err := ioutil.ReadFile(filename)
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
	out = out[:len(out)-2] + "}"
	fmt.Println(out)

	if c.GlobalBool("clipboard") {
		clipboard.WriteAll(out)
	}

	return
}
