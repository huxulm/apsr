package main

// 5
// 1 2 -3 3 3
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func solve(_r io.Reader, _w io.Writer) {
	r, w := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer w.Flush()
	var n int
	fmt.Fscan(r, &n)
	var a, b, ans []int
	var v int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v)
		if v < 0 {
			a = append(a, v)
		} else if v == 0 {
			b = append(b, v)
		} else {
			ans = append(ans, v)
		}
	}

	sort.Ints(a)

	// <= 0
	if len(ans) == 0 && len(a) < 2 {
		if len(b) > 0 {
			fmt.Fprintln(w, 0)
		} else {
			fmt.Println(a[len(a)-1])
		}
	} else {
		ans = append(ans, a[:len(a)-len(a)%2]...)
	}
	sort.Ints(ans)
	for _, v := range ans {
		fmt.Fprintf(w, "%d ", v)
	}
	fmt.Fprintln(w)
}

func main() {
	// in := strings.NewReader(`5
	// 1 2 -3 3 3`)
	// in := strings.NewReader(`13
	// 100 100 100 100 100 100 100 100 100 100 100 100 100`)
	in := strings.NewReader(`4
	-2 -2 -2 -2`)
	// solve(os.Stdin, os.Stdout)
	solve(in, os.Stdout)
}
