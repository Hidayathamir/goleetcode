package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	assert.Equal(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func TestB(t *testing.T) {
	assert.Equal(t, []int{1}, topKFrequent([]int{1}, 1))
}
