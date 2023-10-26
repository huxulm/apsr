package main

import (
	"testing"
	"strings"
	"bytes"
	"fmt"
)
func TestXxx(t *testing.T) {
	cases := []struct {
		raw string
		expect string
	}{
		{
			`4
			2 4
			2 2
			3 3
			1 5`,
			`3
			2
			-3
			-3`,
		},
		{
			`1
			1 1000000000`,
			`500000000`,
		},
	}
	for i, cas := range cases {
		t.Run(fmt.Sprintf("#test-%d", i + 1), func(t *testing.T) {
			r := strings.NewReader(cas.raw)
			w := &bytes.Buffer{}
			_ = w
			solve(r, w)
			t.Log(w.String())			
		})
	}
}