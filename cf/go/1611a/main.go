package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
)
// 输入
// 4
// 3876
// 387
// 4489
// 3
// 输出
// 0
// 2
// 1
// -1

func solve(_r io.Reader, _w io.Writer) {
	in, out := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		nums := make([]int, len(s))
		for i := range s {
			nums[i] = int(s[i] & 15)
		}
		n := len(nums)
		if nums[n-1] & 1 == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		if nums[0] & 1 == 0 {
			fmt.Fprintln(out, 1)
			continue
		}
		ok := false
		for _, v := range nums {
			if v & 1 == 0 {
				ok = true
			}
		}
		if ok {
			fmt.Fprintln(out, 2)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
func main() {
	solve(os.Stdin, os.Stdout)
}