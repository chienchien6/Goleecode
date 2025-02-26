package main

import "fmt"

//理解指針移動：相向雙指針的關鍵是 left 找 val，right 找非 val，然後交換，把 val 推到右邊。

func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	// 印出初始狀態
	fmt.Printf("初始陣列: %v, val: %d, left: %d, right: %d\n", nums, val, left, right)

	for left <= right {
		// 尋找左邊的第一個 val
		for left <= right && nums[left] != val {
			left++
		}
		fmt.Printf("左指針移動後: left=%d, nums=%v\n", left, nums)

		// 尋找右邊的第一個非 val
		for left <= right && nums[right] == val {
			right--
		}
		fmt.Printf("右指針移動後: right=%d, nums=%v\n", right, nums)

		// 交換並移動指針
		if left < right {
			nums[left] = nums[right]
			left++
			right--
			fmt.Printf("交換後: left=%d, right=%d, nums=%v\n", left, right, nums)
		}
	}

	// 印出最終結果
	fmt.Printf("最終結果: left=%d, nums=%v\n", left, nums)
	return left
}

func main() {
	// 測試範例
	nums := []int{3, 2, 2, 3}
	val := 3
	result := removeElement(nums, val)
	fmt.Printf("新長度: %d, 前 %d 個元素: %v\n", result, result, nums[:result])
}
