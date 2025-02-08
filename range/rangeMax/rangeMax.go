package main

import "fmt"

func main() {
	numbers := []int{5, 2, 8, 1, 4, 9}
	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	fmt.Println("最大值:", max)
}
