// https://www.luogu.com.cn/problem/CF1251B
// https://www.luogu.com.cn/record/140939653
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(_r io.Reader, _w io.Writer) {
	in, out := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer out.Flush()
	var t, n int
	var s string
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		var hasOdd bool
		c1 := 0
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &s)
			if len(s)%2 > 0 {
				hasOdd = true
			}
			c1 += strings.Count(s, "1")
		}
		ans := n
		if !hasOdd && c1%2 > 0 {
			ans--
		}
		fmt.Fprintln(out, ans)
	}
}

func main() { solve(os.Stdin, os.Stdout) }
