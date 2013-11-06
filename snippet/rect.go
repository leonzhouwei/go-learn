package main

import "fmt"

type Rect struct {
    x, y float64
    width, height float64
}

func (r *Rect) Area() float64 {
    return r.width * r.height;
}

func main() {
    r1 := new(Rect);
    fmt.Println(r1, r1.Area())
    r2 := &Rect{}
    fmt.Println(r2, r2.Area())
    r3 := &Rect{0, 0, 100, 200}
    fmt.Println(r3, r3.Area())
    r4 := &Rect{width: 300, height: 500}
    fmt.Println(r4, r4.Area())
}
