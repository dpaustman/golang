package main

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Value int
}

func (root *TreeNode) traverse() {
	if root == nil {
		return
	}
	root.Left.traverse()
	fmt.Println(root.Value)
	root.Right.traverse()
	fmt.Println(root.Value)
}

func main() {
	
}
