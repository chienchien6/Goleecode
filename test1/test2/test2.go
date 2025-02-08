package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]

	c := make([]string, 4)
	copy(c, b)
	copy(c[2:], a)
	fmt.Println(a, b)
	fmt.Println("c:", c)
	fmt.Printf("%T\n", c)
	fmt.Printf("c:%#v\n", c)
	fmt.Printf("c:%v\n", c)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// >>
// [John Paul George Ringo]
// [John Paul] [Paul George]
// [John XXX] [XXX George]
// [John XXX George Ringo]
