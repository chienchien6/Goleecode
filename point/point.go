package main

import "fmt"

func increment(n *int) {
	*n += 10
}

func main() {
	num := 5
	increment(&num)
	fmt.Println(num) // 輸出: 15
}

// 這個例子展示了如何使用指針在函數中修改變量的值。在 increment 函數中，我們將傳入的指針參數 n 所指向的變量的值增加 10。在 main 函數中，我們創建了一個整數變量 num，並將其地址傳遞給 increment 函數。由於 increment 函數接受的是 num 的地址，所以在函數內部對 n 的操作會影響到 main 函數中的 num 變量。因此，最終輸出的結果是 15。
