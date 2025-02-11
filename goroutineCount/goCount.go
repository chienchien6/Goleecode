package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	var wg sync.WaitGroup
	ch := make(chan int)

	//啟動多個 goroutine
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}

	//另一個 goroutine 監聽通道
	go func() {
		for range ch {
			count++
			if count == 6 {
				close(ch)
			}
		}
	}()

	wg.Wait()
	fmt.Printf("goroutine count: %d\n", count)
}
