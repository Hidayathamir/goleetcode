package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4})

	reorderList(head)

	assert.Equal(t, []int{1, 4, 2, 3}, head.GetVals())
}

func TestB(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5})

	reorderList(head)

	assert.Equal(t, []int{1, 5, 2, 4, 3}, head.GetVals())
}
