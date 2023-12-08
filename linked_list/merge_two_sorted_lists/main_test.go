package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	list1 := NewListNode([]int{1, 2, 4})
	list2 := NewListNode([]int{1, 3, 4})

	head := mergeTwoLists(list1, list2)

	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, head.GetVals())
}

func TestB(t *testing.T) {
	list1 := NewListNode([]int{})
	list2 := NewListNode([]int{})

	head := mergeTwoLists(list1, list2)

	assert.Equal(t, []int(nil), head.GetVals())
}

func TestC(t *testing.T) {
	list1 := NewListNode([]int{})
	list2 := NewListNode([]int{0})

	head := mergeTwoLists(list1, list2)

	assert.Equal(t, []int{0}, head.GetVals())
}

func TestD(t *testing.T) {
	list1 := NewListNode([]int{1, 2, 5})
	list2 := NewListNode([]int{1, 3, 6})

	head := mergeTwoLists(list1, list2)

	assert.Equal(t, []int{1, 1, 2, 3, 5, 6}, head.GetVals())
}
