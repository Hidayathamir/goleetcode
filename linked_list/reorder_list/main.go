package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	arrayNode := []*ListNode{}

	for head != nil {
		arrayNode = append(arrayNode, head)
		head = head.Next
	}

	l := 0
	r := len(arrayNode) - 1
	p := -1

	for l <= r {
		if p != -1 {
			arrayNode[p].Next = arrayNode[l]
		}

		if l != r {
			arrayNode[l].Next = arrayNode[r]
		}
		p = r

		l++
		r--
	}

	arrayNode[p].Next = nil
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
