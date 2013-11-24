package demopkg

type Rect struct {
    x, y int
    Width, Height float32
}

func (r *Rect) Area() float32 {
    return r.Width * r.Height
}
