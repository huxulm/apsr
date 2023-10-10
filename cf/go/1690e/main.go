package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func solve(_in io.Reader, _out io.Writer) {
	in, out := bufio.NewReader(_in), bufio.NewWriter(_out)
	defer out.Flush()
	var t, n, k int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &k)
		a := make([]int, n)
		ans := 0
		for j := range a {
			fmt.Fscan(in, &a[j])
			ans += a[j] / k
			a[j] %= k
		}
		sort.Ints(a)
		s, e := 0, n-1
		for s < e {
			if a[s]+a[e] < k {
				s++
			} else {
				ans += 1
				s++
				e--
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
