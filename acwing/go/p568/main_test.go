package main

import (
	"testing"
	"strings"
	"bytes"
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
	}
	for _, cas := range cases {
		r := strings.NewReader(cas.raw)
		w := &bytes.Buffer{}
		solve(r, w)
		if w.String() == cas.expect {
			t.Log("ok")
		}
	}
}