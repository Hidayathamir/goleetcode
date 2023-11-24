package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	actual := productExceptSelf([]int{1, 2, 3, 4})
	expected := []int{24, 12, 8, 6}
	assert.Equal(t, expected, actual)
}

func TestMainB(t *testing.T) {
	actual := productExceptSelf([]int{-1, 1, 0, -3, 3})
	expected := []int{0, 0, 9, 0, 0}
	assert.Equal(t, expected, actual)
}
