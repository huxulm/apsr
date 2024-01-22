// https://www.luogu.com.cn/problem/CF1748B
package main

import (
	"fmt"
)

func main() {
	var (
		t, n int
		s    string
	)

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n, &s)
		var ans int
		for i := range s {
			var k, mxC int
			var cnt [10]int
			for _, v := range s[i:] {
				v -= '0'
				if cnt[v] == 10 {
					break
				}
				if cnt[v] == 0 {
					k++
				}
				cnt[v]++
				mxC = max(mxC, cnt[v])
				if mxC <= k {
					ans++
				}
			}
		}
		fmt.Println(ans)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
