package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, l, r, ans int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &l, &r)
		k := (r - l + 1) >> 1
		if l & 1 == 1 {
			ans = k
		} else {
			ans = -k
		}
		if (r - l + 1) & 1 == 1 {
			if r & 1 == 1 {
				ans -= r 
			}
		}
		fmt.Fprint(out, ans)
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}