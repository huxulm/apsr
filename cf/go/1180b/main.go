package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, t int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] >= 0 {
			a[i] = -a[i] - 1
		}
		if a[i] < a[t] {
			t = i
		}
	}
	if n%2 > 0 {
		a[t] = -a[t] - 1
	}
	for _, v := range a {
		fmt.Fprint(out, v, " ")
	}
}
func main() {
	solve(os.Stdin, os.Stdout)
}
