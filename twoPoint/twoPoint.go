package main

func twoSum(pages []int, target int) []int {
	left, right := 0, len(pages)-1
	for left < right {
		sum := pages[left] + pages[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	pages := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(pages, target)
	if result != nil {
		println(result[0], result[1])
	}
}
