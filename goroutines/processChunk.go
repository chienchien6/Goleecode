package main

import (
	"fmt"
	"sync"
)

func processChunk(chunk []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 處理數據邏輯
	for _, v := range chunk {
		// 假設進行某種計算
		fmt.Println(v * 2)
	}
}

func main() {
	data := make([]int, 10000000) // 模擬 1000 萬筆數據
	for i := 0; i < 10000000; i++ {
		data[i] = i
	}

	chunkSize := 100000 // 每個塊的大小
	var wg sync.WaitGroup

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go processChunk(data[i:end], &wg)
	}

	wg.Wait() // 等待所有 goroutines 完成
}
