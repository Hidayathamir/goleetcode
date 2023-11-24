package main

func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i, num := range nums {
		if iHashmap, ok := hashmap[num]; ok {
			return []int{iHashmap, i}
		}
		hashmap[target-num] = i
	}
	return nil
}
