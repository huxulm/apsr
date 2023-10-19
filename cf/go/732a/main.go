package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	in, out := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer out.Flush()

	var k, r int
	fmt.Fscan(in, &k, &r)
	i := 1
	for {
		if ((k % 10) * i) % 10 == 0 || ((k % 10) * i) % 10 == r {
			break
		}
		i++
	}
	fmt.Fprintln(out, i)
}

func main() { solve(os.Stdin, os.Stdout) }