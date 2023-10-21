package main

import(
    "bufio"
    "io"
    "fmt"
    "os"
)

func solve(_r io.Reader, _w io.Writer) {
    in, out := bufio.NewReader(_r), bufio.NewWriter(_w)
    defer out.Flush()
    
    var t int
    fmt.Fscan(in, &t)
    
    for i := 0; i < t; i++ {
        var l, r int
        fmt.Fscan(in, &l, &r)
        if r / l < 2 {
            fmt.Fprintln(out, -1, -1)
        } else {
            fmt.Fprintln(out, l, l *2)
        }
    }
}

func main() { solve(os.Stdin, os.Stdout) }