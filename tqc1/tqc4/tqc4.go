package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	var n int

	// 輸入兩個字串和正整數 n
	fmt.Print("請輸入第一個字串: ")
	fmt.Scanln(&str1)
	fmt.Print("請輸入第二個字串: ")
	fmt.Scanln(&str2)
	fmt.Print("請輸入正整數 n: ")
	fmt.Scanln(&n)

	// 檢查字串長度是否相同
	if len(str1) != len(str2) {
		fmt.Println("error")
		return
	}

	// 檢查字串長度是否超過 128 個字元
	if len(str1) > 128 || len(str2) > 128 {
		fmt.Println("error")
		return
	}

	// 比較前 n 個字元
	if len(str1) < n || len(str2) < n {
		fmt.Println("error")
		return
	}

	cmp := strings.Compare(str1[:n], str2[:n])
	if cmp < 0 {
		fmt.Println("小於")
	} else if cmp > 0 {
		fmt.Println("大於")
	} else {
		fmt.Println("等於")
	}

}
