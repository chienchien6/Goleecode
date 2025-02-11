package main

import "fmt"

// /定義一個接口 Shape，包含 Area() 和 Perimeter() 方法，然後實現該接口的 Circle 和 Rectangle 結構體。
type Shape interface {
	Area() float64      // 面積
	Perimeter() float64 // 周長
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64 //半徑
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {

	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	cArea := c.Area()
	cPerimeter := c.Perimeter()
	rArea := r.Area()
	rPerimeter := r.Perimeter()
	fmt.Printf("Circle Area: %f, Circle Perimeter: %f\n", cArea, cPerimeter)
	fmt.Printf("Rectangle Area: %f, Rectangle Perimeter: %f\n", rArea, rPerimeter)

}
