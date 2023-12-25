package main

import (
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	in := strings.NewReader(`4
	1
	0
	3
	1110
	100110
	010101
	2
	11111
	000001
	2
	001
	11100111`)
	solve(in, os.Stdout)
}
