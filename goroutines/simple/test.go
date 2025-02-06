package main

import (
	"fmt"
	"sync"
)

func calculateSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done()                                        // 當計算完成時，減少等待計數
	square := number * number                              // 計算平方
	fmt.Printf("The square of %d is %d\n", number, square) // 打印結果
}

func main() {
	numbers := []int{1, 2, 3, 4, 5} // 一組數字
	var wg sync.WaitGroup           // 等待組

	for _, number := range numbers {
		wg.Add(1)                       // 增加等待計數
		go calculateSquare(number, &wg) // 啟動一個新的 goroutine
	}

	wg.Wait() // 等待所有計算完成
}
