package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func TestMainB(t *testing.T) {
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestMainC(t *testing.T) {
	assert.Equal(t, 2, maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
}

func TestMainD(t *testing.T) {
	assert.Equal(t, 8, maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}
