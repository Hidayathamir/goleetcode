package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tempNode := root.Left

	root.Left = root.Right
	root.Right = tempNode

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
