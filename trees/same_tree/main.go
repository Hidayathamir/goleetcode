package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	isLeftSameTree := isSameTree(p.Left, q.Left)
	isRightSameTree := isSameTree(p.Right, q.Right)

	return isLeftSameTree && isRightSameTree
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
