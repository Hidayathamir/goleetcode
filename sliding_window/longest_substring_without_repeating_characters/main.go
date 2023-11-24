package main

import "strings"

func lengthOfLongestSubstring(s string) int {
	maximumLength := 0
	l := 0
	for r, v := range s {
		for strings.Contains(s[l:r], string(v)) && l < r {
			l++
		}

		currentLength := r - l + 1
		maximumLength = max(maximumLength, currentLength)
	}
	return maximumLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
