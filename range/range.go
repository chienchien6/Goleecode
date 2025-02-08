package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}
}
