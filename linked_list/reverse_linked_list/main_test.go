package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5})

	reversedHead := reverseList(head)

	assert.Equal(t, []int{5, 4, 3, 2, 1}, reversedHead.GetVals())
}

func TestB(t *testing.T) {
	head := NewListNode([]int{1, 2})

	reversedHead := reverseList(head)

	assert.Equal(t, []int{2, 1}, reversedHead.GetVals())
}

func TestC(t *testing.T) {
	head := NewListNode([]int{})

	reversedHead := reverseList(head)

	assert.Equal(t, []int(nil), reversedHead.GetVals())
}
