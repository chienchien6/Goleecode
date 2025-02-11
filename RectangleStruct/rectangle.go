package main

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 12, Height: 2}
	area := r.Area()
	println(area)
}
