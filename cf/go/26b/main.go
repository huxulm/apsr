package main

import (
	"fmt"
	"bufio"
	"io"
	// "os"
)

// Input1
// (()))(
// Ouput1
// 4

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	st := 0
	ans := 0
	for i := range s {
		if s[i] == '(' {
			st++
		} else {
			if st > 0 {
				st--
				ans++
			}
		}
	}
	fmt.Fprintln(out, ans * 2)
}