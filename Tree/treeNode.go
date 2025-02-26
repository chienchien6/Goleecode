package main

import (
	"fmt"
)

// 定義二元搜尋樹的節點
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 插入節點到 BST
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val: val}
	}
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		root.right = insert(root.right, val)
	}
	return root
}

// 前序遍歷 (Preorder Traversal)
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.val, " ") // 先輸出根節點
	preorderTraversal(root.left)
	preorderTraversal(root.right)
}

func main() {
	var n int
	fmt.Scan(&n) // 讀取 N

	var root *TreeNode
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val) // 讀取每個數字
		root = insert(root, val)
	}

	preorderTraversal(root) // 執行前序遍歷
	fmt.Println()           // 換行
}
