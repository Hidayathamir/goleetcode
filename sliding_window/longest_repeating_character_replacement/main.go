package main

func characterReplacement(s string, k int) int {
	count := map[string]int{}
	res := 0

	l := 0
	for r := range s {
		count[string(s[r])]++
		for (r-l+1)-maxHashmap(count) > k {
			count[string(s[l])]--
			l++
		}
		res = maxInt(res, r-l+1)
	}
	return res
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxHashmap(count map[string]int) int {
	max := 0
	for _, sum := range count {
		if sum > max {
			max = sum
		}
	}
	return max
}
