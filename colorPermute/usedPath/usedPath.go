package main

import "fmt"

func findPaths(path []string, used []bool) {
	// 假設有三個路徑
	for i := 0; i < len(path); i++ {
		if !used[i] { // 如果這個路徑沒有被使用過
			used[i] = true // 標記為已使用
			fmt.Println("使用路徑:", path[i])
			findPaths(path, used) // 遞歸調用
			used[i] = false       // 回溯，標記為未使用
		}
	}
}

func main() {
	paths := []string{"A", "B", "C"}
	used := make([]bool, len(paths)) // 初始化 used 切片
	findPaths(paths, used)           // 開始尋找路徑
}
