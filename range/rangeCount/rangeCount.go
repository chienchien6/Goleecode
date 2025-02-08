package main

import "fmt"

func main() {
	str := "hello"
	frequency := make(map[rune]int)
	for _, char := range str {
		frequency[char]++
	}
	for char, count := range frequency {
		fmt.Printf("字符: %c, 頻率: %d\n", char, count)
	}
}
