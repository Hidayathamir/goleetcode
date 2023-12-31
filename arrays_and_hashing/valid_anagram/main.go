package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		hashmap[s[i]]++
		hashmap[t[i]]--
	}

	for _, v := range hashmap {
		if v != 0 {
			return false
		}
	}

	return true
}
