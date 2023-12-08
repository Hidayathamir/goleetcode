package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
