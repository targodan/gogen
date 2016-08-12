package conv

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func BytesToString(data []byte, lnbr int) string {
	out := "[]byte{"
	if lnbr > 0 {
		out = fmt.Sprintln(out)
	}
	ln := 0
	for _, b := range data {
		out = fmt.Sprintf("%s0x%02x, ", out, b)
		ln++
		if lnbr > 0 && ln >= lnbr {
			out = fmt.Sprintln(out)
			ln = 0
		}
	}
	if ln == 0 {
		out = out[:len(out)-1]
	}
	if lnbr > 0 {
		out = fmt.Sprintln(out)
		out += "}"
	}

	return out
}

func readFromStdin() (data []byte, err error) {
	r := bufio.NewReader(os.Stdin)
	data = make([]byte, 0, 16)
	var d byte
	for {
		d, err = r.ReadByte()
		if err != nil {
			break
		}
		data = append(data, d)
	}
	if err == io.EOF {
		err = nil
	}
	data = bytes.Trim(data, " \t\r\n")
	return
}

func FileOrStdin(arg string) (data []byte, err error) {
	if arg != "-" {
		data, err = ioutil.ReadFile(arg)
	} else {
		data, err = readFromStdin()
	}
	return
}

func TextOrStdin(arg string) (string, error) {
	if arg != "-" {
		return arg, nil
	}
	data, err := readFromStdin()
	return string(data), err
}
