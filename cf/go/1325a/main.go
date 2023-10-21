package main

import(
	"fmt"
	"bufio"
	"os"
	"io"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i ++ {
		var x int
		fmt.Fscan(in, &x)
		g := 1
		l := x - g
		fmt.Fprintln(out, g, l)
	}
}

func main() { solve(os.Stdin, os.Stdout) }