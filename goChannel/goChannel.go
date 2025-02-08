package main

import (
	"fmt"
	"time"
)

type Message struct {
	Content string
}

func main() {
	aChannel := make(chan Message, 10) // 緩衝區避免阻塞
	bChannel := make(chan Message, 2)  // B 最多拿 2 顆球
	cChannel := make(chan Message, 2)  // C 最多拿 2 顆球
	msgChannel := make(chan string)

	go messageSender(msgChannel)

	go person("A", aChannel, bChannel, msgChannel)
	go person("B", bChannel, cChannel, msgChannel)
	go person("C", cChannel, nil, msgChannel)

	// 模擬球掉下來
	go func() {
		for {
			time.Sleep(5 * time.Second) // 每 5 秒掉一顆球
			aChannel <- Message{Content: "球"}
		}
	}()

	// 接收訊息
	for {
		select {
		case msg := <-aChannel:
			fmt.Println("A 收到:", msg.Content)
			bChannel <- msg
		case msg := <-bChannel:
			fmt.Println("B 收到:", msg.Content)
			cChannel <- msg
		case msg := <-cChannel:
			fmt.Println("C 收到:", msg.Content)
		case msg := <-msgChannel:
			fmt.Println("訊息發送機:", msg)
		}
	}
}

func person(name string, receive <-chan Message, send chan<- Message, msgChannel chan<- string) {
	ballCount := 0 // 手上的球數

	for {
		select {
		case msg := <-receive:
			fmt.Printf("%s 收到: %s\\n", name, msg.Content)
			ballCount++

			if send != nil {
				// 如果下一個人手上已經有 2 顆球，發送催促訊息
				if len(send) >= 2 {
					msgChannel <- fmt.Sprintf("%s 請快點接球!", name)
				} else {
					send <- msg
					ballCount--
				}
			} else {
				// C 丟球的邏輯
				if ballCount >= 2 {
					fmt.Printf("%s 丟掉一顆球\\n", name)
					time.Sleep(1 * time.Second) // 丟球需要時間
					ballCount--
				}
			}
		default:
			// 如果手上球數達到 2 顆，發送催促訊息
			if ballCount >= 2 && (name == "B" || name == "C") {
				msgChannel <- fmt.Sprintf("%s 請快點接球!", name)
				time.Sleep(2 * time.Second) // 模擬等待時間
			}
		}
	}
}

func messageSender(msgChannel chan string) {
	for {
		msg := <-msgChannel
		fmt.Println("訊息發送機:", msg)
	}
}
