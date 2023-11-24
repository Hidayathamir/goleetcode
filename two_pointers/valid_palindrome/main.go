package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = onlyAlphaNumeric(s)

	if s == "" {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func onlyAlphaNumeric(input string) string {
	regex := regexp.MustCompile("[^a-z0-9]+")
	return regex.ReplaceAllString(input, "")
}
