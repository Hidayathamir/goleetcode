package main

func reverseList(head *ListNode) *ListNode {
	var prevHead *ListNode
	for head != nil {
		currentHead := head
		nextHead := head.Next

		head.Next = prevHead
		head = nextHead
		prevHead = currentHead
	}
	return prevHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	firstNode := head

	for i := 1; i < len(vals); i++ {
		head.Next = &ListNode{Val: vals[i]}
		head = head.Next
	}

	return firstNode
}

func (l *ListNode) GetVals() []int {
	var vals []int

	for l != nil {
		vals = append(vals, l.Val)
		l = l.Next
	}

	return vals
}
