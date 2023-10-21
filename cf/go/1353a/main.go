// https://www.luogu.com.cn/problem/CF1353A
// https://codeforces.com/problemset/submission/1353/228999103
package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
)

func solve(_r io.Reader, _w io.Writer) {
    in := bufio.NewReader(_r)
    out := bufio.NewWriter(_w)
    defer out.Flush()
    var t int
    fmt.Fscan(in, &t)
    for i := 0; i < t; i++ {
        var n, m int 
        fmt.Fscan(in, &n, &m)
        if n == 1 {
            fmt.Fprintln(out, 0)
        } else if n == 2 {
            fmt.Fprintln(out, m)
        } else {
            fmt.Fprintln(out, m * 2)
        }
    }
}

func main() { solve(os.Stdin, os.Stdout) }