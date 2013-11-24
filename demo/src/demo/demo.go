package main

import "fmt"

import "demopkg"

func main() {
    r := &demopkg.Rect{}
    fmt.Println(r)
    demopkg.Hello()
}
