package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
}

func TestMainB(t *testing.T) {
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
}

func TestMainC(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}

func TestMainD(t *testing.T) {
	assert.Equal(t, 2, lengthOfLongestSubstring("abba"))
}

func TestMainE(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
}
