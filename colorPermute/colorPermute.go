package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue"}
	result := permute(colors)
	fmt.Println(result)
}

func permute(colors []string) [][]string {
	var res [][]string
	var backtrack func(path []string, used []bool)
	backtrack = func(path []string, used []bool) {
		if len(path) == len(colors) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := 0; i < len(colors); i++ {
			if used[i] {
				continue // 如果顏色已經使用過，則跳過
			}
			used[i] = true                 // 標記顏色為已使用
			path = append(path, colors[i]) // 將顏色加入當前路徑
			backtrack(path, used)          // 進行回溯
			used[i] = false                // 回溯後，將顏色標記為未使用
			path = path[:len(path)-1]      // 從當前路徑中移除最後一個顏色
		}

	}
	backtrack([]string{}, make([]bool, len(colors)))
	return res
}
