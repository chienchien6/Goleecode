package main

import "fmt"

func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	min, max := minMax(nums)
	fmt.Printf("最小值: %d, 最大值: %d\n", min, max)

	//寫法2
	min, max = nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}

	}
	fmt.Printf("寫法2==最小值: %d, 最大值: %d\n", min, max)

}
