package main

func containsDuplicate(nums []int) bool {
	numsHashmap := map[int]bool{}
	for _, num := range nums {
		if _, ok := numsHashmap[num]; ok {
			return true
		}
		numsHashmap[num] = true
	}
	return false
}
