package main

import (
    "fmt"
    "os"
    "strings"
)

func main(){
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Println(os.Args[1:])
    for x, ag := range os.Args[1:]{
        fmt.Println(x, ag)
    }
}
