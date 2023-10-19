// https://www.luogu.com.cn/problem/CF148A
package main

import(
    "fmt"
    "os"
    "bufio"
    "io"
)

// 输入1
// 1
// 2
// 3
// 4
// 12
// 输出1
// 12

// 输入2
// 2
// 3
// 4
// 5
// 24
// 输出2
// 17
func main() {
    solve(os.Stdin, os.Stdout)
}

func solve(_r io.Reader, _w io.Writer) {
    reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
    defer writer.Flush()
    
    var k, l, m, n int
    var d int
    fmt.Fscan(reader, &k, &l, &m, &n, &d)
    ans := 0
    for i := 1; i < d + 1; i++ {
        if i % k == 0 || i % l == 0 || i % m == 0 || i % n == 0 {
            ans++
        }
    }
    fmt.Fprintln(writer, ans)
}