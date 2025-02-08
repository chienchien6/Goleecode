package main

import "fmt"

//	for {} 就可以進入一個無限迴圈
func main() {
	for {
		var score int

		fmt.Println()
		fmt.Print("輸入分數: ")
		fmt.Scanln(&score)

		// if score < 0 || score > 100 {
		// 	fmt.Println("error")
		// }

		// if score > 60 {
		// 	score += 10
		// 	fmt.Printf("調整後的分數:%#v", score)

		// } else {
		// 	score += 5
		// 	fmt.Printf("調整後的分數:%#v", score)
		// }

		if score < 0 || score > 100 {
			fmt.Println("error")
		} else if score > 60 {
			score += 10
			fmt.Printf("調整後的分數:%#v", score)

		} else {
			score += 5
			fmt.Printf("調整後的分數:%#v", score)
		}
	}
}
