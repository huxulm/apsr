package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func solve(_r io.Reader, _w io.Writer) {
	in, out := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer out.Flush()

	s, _, _ := in.ReadLine()
	mp := map[byte]bool{}
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			mp[s[i]] = true
		}
	}
	fmt.Fprintln(out, len(mp))
}

func main() {
	solve(os.Stdin, os.Stdout)
}