package main

import (
	"fmt"
)

// 計算n的階乘
func factorial(n int) int {
	// 基本情況
	if n == 0 {
		return 1
	}
	// 遞回調用
	return n * factorial(n-1)
}

func main() {
	num := 5
	result := factorial(num)
	fmt.Printf("%d 的階乘是: %d\n", num, result)
}
