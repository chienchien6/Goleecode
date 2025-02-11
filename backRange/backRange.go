package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Printf("%#v", nums[i])
	}
}
