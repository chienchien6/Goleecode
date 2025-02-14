package binarySearch

// 时间复杂度 O(logn)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // 這裡要減 1，因為陣列是從 0 開始的

	// 循环逐步缩小区间范围
	for left <= right {

		mid := left + (right-left)>>1 // >> 1右移一位，相當於除以 2，這樣可以獲得範圍的中點。避免溢位 =(right + left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1

}

// 时间复杂度 O(logn)
func searchOpen(nums []int, target int) int {
	left := 0
	right := len(nums) // 這裡不用減 1，因為這裡是開區間

	// 循环逐步缩小区间范围
	for left < right {

		mid := left + (right-left)>>1 // >> 1右移一位，相當於除以 2，這樣可以獲得範圍的中點。避免溢位 =(right + left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
