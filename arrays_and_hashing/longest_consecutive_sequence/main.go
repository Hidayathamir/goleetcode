package main

func longestConsecutive(nums []int) int {
	hashmap := map[int]int{}
	for i, number := range nums {
		hashmap[number] = i
	}

	maxCount := 0
	for _, number := range nums {
		if _, ok := hashmap[number-1]; ok {
			continue
		}

		count := 0
		for {
			if _, ok := hashmap[number]; ok {
				count++
				number++
			} else {
				break
			}
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
