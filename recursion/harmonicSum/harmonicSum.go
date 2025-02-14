package main

import (
	"fmt"
)

// 遞迴計算 Harmonic Sum
func harmonicSum(n int) float64 {
	if n <= 0 {
		return 0.0
	}
	if n == 1 {
		return 1.0
	}
	return 1.0/float64(n) + harmonicSum(n-1)
}

// 當基本條件滿足時，遞迴不會進入無窮循環，而是開始回溯計算結果。
func main() {
	var n int
	fmt.Print("請輸入正整數 n: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("請輸入正整數")
		return
	}

	fmt.Printf("Harmonic Sum H(%d) = %.6f\n", n, harmonicSum(n))
}
