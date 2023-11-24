package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	actual := twoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	assert.Equal(t, expected, actual)
}

func TestMainB(t *testing.T) {
	actual := twoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	assert.Equal(t, expected, actual)
}

func TestMainC(t *testing.T) {
	actual := twoSum([]int{3, 3}, 6)
	expected := []int{0, 1}
	assert.Equal(t, expected, actual)
}
