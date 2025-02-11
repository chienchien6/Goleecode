package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func updateP(p *Person) {
	p.Name = "Tom"
	p.Age = 31
}

func main() {
	person := Person{
		Name: "John",
		Age:  30,
	}
	updateP(&person)
	fmt.Println(person.Name, person.Age)

}
