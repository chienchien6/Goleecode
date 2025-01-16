package main

import "fmt"

func main() {
	nums := []int{2, 3, 6, 7, 11, 15}
	target := 9
	fmt.Println(MaptwoSum(nums, target))

	fmt.Println(SumtwoSum(nums, target))
}

func MaptwoSum(nums []int, target int) [][]int {
	m := make(map[int]int)
	var result [][]int

	for i, v := range nums { //是當前數字 v 在 nums 切片中的索引。
		if j, ok := m[target-v]; ok { //j：代表在映射 m 中找到的對應索引
			result = append(result, []int{j, i})
		}
		m[v] = i
	}
	return result
}

func SumtwoSum(nums []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
