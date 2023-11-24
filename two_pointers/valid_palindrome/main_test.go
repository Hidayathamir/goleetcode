package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainA(t *testing.T) {
	assert.True(t, isPalindrome("A man, a plan, a canal: Panama"))
}

func TestMainB(t *testing.T) {
	assert.False(t, isPalindrome("race a car"))
}

func TestMainC(t *testing.T) {
	assert.True(t, isPalindrome(" "))
}

func TestMainD(t *testing.T) {
	assert.False(t, isPalindrome("0P"))
}
