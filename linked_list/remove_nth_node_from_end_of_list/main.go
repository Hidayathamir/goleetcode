package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil && n >= 1 {
		return nil
	}

	firstNode := head

	arrayNode := []*ListNode{}
	for head != nil {
		arrayNode = append(arrayNode, head)
		head = head.Next
	}

	indexWillBeRemoved := len(arrayNode) - n
	if indexWillBeRemoved < 0 {
		return head
	}

	if indexWillBeRemoved == 0 {
		return firstNode.Next
	}

	prevIndex := indexWillBeRemoved - 1
	nextIndex := indexWillBeRemoved + 1

	lastArrayNodeIndex := len(arrayNode) - 1

	if nextIndex > lastArrayNodeIndex {
		arrayNode[prevIndex].Next = nil
	} else {
		arrayNode[prevIndex].Next = arrayNode[nextIndex]
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
