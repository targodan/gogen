package main

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	tmp := os.Stdout
	os.Stdout = nil
	main()
	os.Stdout = tmp
}
