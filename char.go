package main

import "fmt"

func main() {
    s := "Hello, world."
    var c byte = s[0]
    fmt.Println(c == 'H')
    for _, ch := range s {
        fmt.Printf("%c\n", ch) 
    }
}
