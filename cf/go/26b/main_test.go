package main

import (
	"testing"
	"os"
	"strings"
)

func TestXxx(t *testing.T) {
	var inputs = []string{
		"(()))(",
		"((()())",
	}
	for _, raw := range inputs {
		solve(strings.NewReader(raw), os.Stdout)
	}
}