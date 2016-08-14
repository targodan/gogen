# gogen
A handy tool to generate small pieces of go code.

# Installing
```bash
$ go get github.com/targodan/gogen
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

# License
This tool is under the GNU General Public License Version 3.
