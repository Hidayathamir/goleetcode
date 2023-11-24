package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	max := 0
	l, r := 0, len(height)-1
	for l < r {
		lValue := height[l]
		rValue := height[r]
		width := r - l
		area := width * min(lValue, rValue)
		if area > max {
			max = area
		}

		if lValue < rValue {
			l++
		} else {
			r--
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
