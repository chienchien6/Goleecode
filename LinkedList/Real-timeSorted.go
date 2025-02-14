package main

import (
	"fmt"
)

// 定義鏈結串列節點
type Node struct {
	data int
	next *Node
}

// 插入節點到排序後的鏈結串列
func insertSorted(head **Node, value int) {
	newNode := &Node{data: value}

	// 如果鏈結串列是空的，或新的值應該成為新的頭節點
	if *head == nil || (*head).data >= value {
		newNode.next = *head
		*head = newNode
		return
	}

	// 遍歷找到適當的插入位置
	current := *head
	for current.next != nil && current.next.data < value {
		current = current.next
	}

	// 插入新節點
	newNode.next = current.next
	current.next = newNode
}

// 印出鏈結串列
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " -> ")
		head = head.next
	}
	fmt.Println("nil")
}

func main() {
	var head *Node // 初始化空鏈結串列

	for {
		var num int
		fmt.Print("請輸入數字（輸入 -1 結束）：")
		fmt.Scan(&num)

		if num == -1 {
			break
		}

		insertSorted(&head, num)
		printList(head)
	}
}
