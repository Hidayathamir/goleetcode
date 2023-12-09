package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	slowPointer := head
	fastPointer := head.Next

	for slowPointer != fastPointer {
		if slowPointer.Next == nil {
			return false
		}
		slowPointer = slowPointer.Next

		if fastPointer.Next == nil || fastPointer.Next.Next == nil {
			return false
		}
		fastPointer = fastPointer.Next.Next
	}

	return true
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
