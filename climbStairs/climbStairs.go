package main

import "fmt"

func main() {

	count := climbStairs(5)
	fmt.Println(count)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n // 如果樓梯階數小於等於2，直接返回階數
	}
	dp := make([]int, n+1) // 創建一個大小為 n+1 的切片來存儲每階的方式數
	dp[1], dp[2] = 1, 2    // 初始化第1階和第2階的方式數
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 根據前兩階的方式數計算當前階的方式數
	}
	return dp[n] // 返回到達第 n 階的方式數
}
