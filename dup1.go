package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
        if input.Text() == "EOF" {
            break
        }
    }

    for line, n := range counts {
        fmt.Printf("%s\t%d\n", line, n)
//        if n > 1 {
//            fmt.Println("%d\t%s\n", n, line)
//        }
    }
}
