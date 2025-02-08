package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("sum:", sum)
}
