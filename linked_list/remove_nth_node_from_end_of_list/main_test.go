package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	head := removeNthFromEnd(NewListNode([]int{1, 2, 3, 4, 5}), 2)
	assert.Equal(t, []int{1, 2, 3, 5}, head.GetVals())
}

func TestB(t *testing.T) {
	head := removeNthFromEnd(NewListNode([]int{1}), 1)
	assert.Equal(t, []int(nil), head.GetVals())
}

func TestC(t *testing.T) {
	head := removeNthFromEnd(NewListNode([]int{1, 2}), 1)
	assert.Equal(t, []int{1}, head.GetVals())
}

func TestD(t *testing.T) {
	head := removeNthFromEnd(NewListNode([]int{1, 2}), 2)
	assert.Equal(t, []int{2}, head.GetVals())
}
