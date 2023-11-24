package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
}

func TestMainB(t *testing.T) {
	assert.False(t, isAnagram("rat", "car"))
}
