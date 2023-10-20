package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()
	var t int 
	fmt.Fscan(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
	
		// [1, n], n >= 2
		a := make([]int, n)
		for i := 0; i < n; i++ {
			a[i] = n - i
		}
	
		if n & 1 == 1{
			a[0], a[n/2] = a[n/2], a[0] 
		}

		for _, v := range a {
			fmt.Fprintf(writer, "%d ", v)
		}

		fmt.Fprintln(writer)
	}
}