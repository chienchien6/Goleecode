package main

import "fmt"

func main() {

	var midscore int
	fmt.Printf("輸入分數:")
	fmt.Scanln(&midscore)

	adjustScore := compute(midscore)

	fmt.Println(adjustScore)

}

func compute(score int) int {

	if score < 0 || score > 100 {
		return -1
	} else if score >= 60 {
		score += 5
		return score
	} else {
		score += 10
		return score
	}

}
