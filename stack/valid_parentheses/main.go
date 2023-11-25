package main

func isValid(s string) bool {
	hashmap := map[rune]int{}

	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			hashmap[v]++
		} else {
			if v == ']' {
				hashmap['[']--
			} else if v == ')' {
				hashmap['(']--
			} else {
				hashmap['{']--
			}
		}
	}

	for _, v := range hashmap {
		if v < 0 {
			return false
		}
	}

	return true
}
