package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	students := []Student{
		{"John", 21},
		{"Smith", 22},
	}

	for _, student := range students {
		fmt.Println(student.Name, student.Age)
	}
}
