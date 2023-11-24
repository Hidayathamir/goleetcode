package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	nums := []int{1, 2, 3, 1}

	assert.True(t, containsDuplicate(nums))
}

func TestMainB(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	assert.False(t, containsDuplicate(nums))
}

func TestMainC(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	assert.True(t, containsDuplicate(nums))
}
