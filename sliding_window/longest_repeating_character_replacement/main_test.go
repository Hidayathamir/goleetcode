package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	assert.Equal(t, 4, characterReplacement("ABAB", 2))
}

func TestB(t *testing.T) {
	assert.Equal(t, 4, characterReplacement("AABABBA", 1))
}
