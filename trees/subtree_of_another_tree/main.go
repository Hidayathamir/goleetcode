package main

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	if root.Left != nil {
		if isSubtree(root.Left, subRoot) {
			return true
		}
	}

	if root.Right != nil {
		if isSubtree(root.Right, subRoot) {
			return true
		}
	}

	return false
}

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
