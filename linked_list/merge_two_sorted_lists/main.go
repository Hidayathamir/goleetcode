package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list1 == nil && list2 != nil {
		return list2
	}

	l, r := list1, list2

	var firstNode *ListNode
	var head *ListNode

	if r.Val >= l.Val {
		head = l
		firstNode = head
		l = l.Next
	} else {
		head = r
		firstNode = head
		r = r.Next
	}

	for l != nil && r != nil {
		if r.Val >= l.Val {
			head.Next = l
			l = l.Next
		} else {
			head.Next = r
			r = r.Next
		}

		head = head.Next
	}

	if l != nil {
		head.Next = l
	} else {
		head.Next = r
	}

	return firstNode
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
