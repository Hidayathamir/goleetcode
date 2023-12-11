package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepthLeft := 1 + maxDepth(root.Left)
	maxDepthRight := 1 + maxDepth(root.Right)

	return max(maxDepthLeft, maxDepthRight)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
