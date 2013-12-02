package main

import "fmt"

type Rect struct {
    x, y float64
    width, height float64
}

func NewRect(x, y, width, height float64) *Rect {
    return &Rect{x, y, width, height}
}

func (r *Rect) Area() float64 {
    return r.width * r.height 
}

func main() {
    rect1 := new(Rect)
    fmt.Println(rect1)
    rect2 := &Rect{}
    fmt.Println(rect2)
    rect3 := &Rect{0, 0, 100, 200}
    fmt.Println(rect3)
    rect4 := &Rect{width: 100, height: 200}
    fmt.Println(rect4)
    rect5 := NewRect(1, 2, 3, 4)
    fmt.Println(rect5)
}
