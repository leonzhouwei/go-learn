package main

import "fmt"

import "demopkg"
import "demopkg/subpkg"

func main() {
    r := &demopkg.Rect{}
    fmt.Println(r)
    demopkg.Hello()
    subpkg.Hello()
}
