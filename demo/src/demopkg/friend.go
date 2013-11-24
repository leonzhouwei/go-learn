package demopkg

import "fmt"

func Hello() {
    r := &Rect{1, 1, 1.2, 3.0}
    fmt.Println(r.x)
}
