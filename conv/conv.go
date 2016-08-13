package conv

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	reOuter *regexp.Regexp
	reLit   *regexp.Regexp
	reBin   *regexp.Regexp
	reOct   *regexp.Regexp
	reHex   *regexp.Regexp
	reDec   *regexp.Regexp
)

func init() {
	reOuter = regexp.MustCompile(`(\s*\[\d*\]byte\s*\{)?([^}]*)\}?`)
	reLit = regexp.MustCompile(`^-?(0b[01]+|0x[a-fA-F0-9]+|\d+)`)
	reBin = regexp.MustCompile(`^(-?)0b([01]+)$`)
	reOct = regexp.MustCompile(`^(-?)0(\d+)$`)
	reHex = regexp.MustCompile(`^(-?)0x([a-fA-F0-9]+)$`)
	reDec = regexp.MustCompile(`^(-?)([1-9]\d*|0)$`)
}

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

func litToByteUnsigned(lit string) (byte, error) {
	var b uint64
	var reRes []string
	var err error

	reRes = reBin.FindStringSubmatch(lit)
	if reRes != nil {
		b, err = strconv.ParseUint(reRes[1]+reRes[2], 2, 8)
	} else {
		reRes = reOct.FindStringSubmatch(lit)
		if reRes != nil {
			b, err = strconv.ParseUint(reRes[1]+reRes[2], 8, 8)
		} else {
			reRes = reHex.FindStringSubmatch(lit)
			if reRes != nil {
				b, err = strconv.ParseUint(reRes[1]+reRes[2], 16, 8)
			} else {
				reRes = reDec.FindStringSubmatch(lit)
				if reRes != nil {
					b, err = strconv.ParseUint(reRes[1]+reRes[2], 10, 8)
				} else {
					return 0, errors.New("Invalid literal. \"" + lit + "\"")
				}
			}
		}
	}

	return byte(b), err
}

func litToByteSigned(lit string) (byte, error) {
	var b int64
	var reRes []string
	var err error

	reRes = reBin.FindStringSubmatch(lit)
	if reRes != nil {
		b, err = strconv.ParseInt(reRes[1]+reRes[2], 2, 8)
	} else {
		reRes = reOct.FindStringSubmatch(lit)
		if reRes != nil {
			b, err = strconv.ParseInt(reRes[1]+reRes[2], 8, 8)
		} else {
			reRes = reHex.FindStringSubmatch(lit)
			if reRes != nil {
				b, err = strconv.ParseInt(reRes[1]+reRes[2], 16, 8)
			} else {
				reRes = reDec.FindStringSubmatch(lit)
				if reRes != nil {
					b, err = strconv.ParseInt(reRes[1]+reRes[2], 10, 8)
				} else {
					return 0, errors.New("Invalid literal. \"" + lit + "\"")
				}
			}
		}
	}

	return byte(b), err
}

func TextToByteSlice(text string) ([]byte, error) {
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	// trims optional "[]byte{...}" wrapping
	byteString := reOuter.FindStringSubmatch(text)[2]
	byteString = strings.Trim(byteString, " \t\r\n,")

	ret := make([]byte, 0, len(byteString)/6)
	var reRes []string
	var b byte
	var err error
	for len(byteString) > 0 {
		reRes = reLit.FindStringSubmatch(byteString)
		if reRes == nil {
			return nil, errors.New("Invalid input string.")
		}
		b, err = litToByteSigned(reRes[0])
		if err != nil {
			b, err = litToByteUnsigned(reRes[0])
			if err != nil {
				return nil, err
			}
		}
		ret = append(ret, b)

		byteString = strings.Trim(byteString[len(reRes[0]):], " \t\r\n,")
	}
	return ret, nil
}
