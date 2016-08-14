# gogen [![Build Status](https://travis-ci.org/targodan/gogen.svg?branch=master)](https://travis-ci.org/targodan/gogen)
A handy tool to generate small pieces of go code.

# Installing
You need at go version 1.4.
This tool is tested against 1.4, 1.5 and 1.6.
If you don't have it yet, get go [here](https://golang.org/).

Once you have go, all you need to do in order to install or update gogen is:

```bash
$ go get -u github.com/targodan/gogen
```

# Basic usage
Base64 en- and decoding.
```
$ gogen base64 encode 0x01,0xff,-64,3,0b10011
Af/AAxM=
```

```
$ gogen base64 decode Af/AAxM=
[]byte{
0x01, 0xff, 0xc0, 0x03, 0x13,
}
```

Hex en- and decoding.

```
$ gogen hex encode 0x01,0xff,-64,3,0b10011
01ffc00313
```

```
$ gogen hex decode 01ffc00313
[]byte{
0x01, 0xff, 0xc0, 0x03, 0x13,
}
```

Embedding files in code.

```
$ gogen file2bytes helloWorld.txt
[]byte{
0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
}
```

# Modifying behaviour and output

All commands support the parameters `--clipboard, -c` and `--linebreak n, -b n`.

The `--clipboard, -c` parameter makes the command copy all output to the clipboard.

The `--linebreak n, -b n` parameter makes the command break each line after n bytes/characters.

```
$ gogen f2b -b 4 test.txt
[]byte{
0x30, 0x78, 0x34, 0x38,
0x2c, 0x20, 0x30, 0x78,
0x36, 0x35, 0x2c, 0x20,
0x30, 0x78, 0x36, 0x63,
0x2c, 0x20, 0x30, 0x78,
0x36, 0x63, 0x2c, 0x20,
0x30, 0x78, 0x36, 0x66,
0x2c, 0x20, 0x30, 0x78,
0x32, 0x30, 0x2c,
}
```

# Contributing

If you want to contribute feel free to create issues or to fork and send a pull request.
Pull requests only on the develop branch and `gofmt` your code prior to committing please.

If you want to add a new command please create a new folder with the same name as your command
and add an import line in the `gogen.go` file of the root directory like so.

```go
import (
    ...
	_ "github.com/targodan/gogen/YOUR_CMD_NAME"
    ...
)
```

Any new files you add need to include the licenseheader (can be found in `licenseheader.txt`)-
Please also add your name to the copyright section of any file you modify like so.

```go
/*
 * Copyright (C) 2016 Some Name, Some Other Name, Even More Names,
 *                    YOUR NAME
 * ...
 */
```

The order of contributers here is first-come-first-serve.

# License
This tool is under the GNU General Public License Version 3.
