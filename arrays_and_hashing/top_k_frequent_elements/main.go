package main

func topKFrequent(nums []int, k int) []int {
	hashmap := map[int]int{}

	for _, num := range nums {
		hashmap[num]++
	}

	list := make([][]int, len(nums)+1)
	for num, sum := range hashmap {
		list[sum] = append(list[sum], num)
	}

	output := []int{}

	for i := len(list) - 1; i >= 0; i-- {
		x := list[i]
		for _, v := range x {
			output = append(output, v)
			if len(output) == k {
				return output
			}
		}
	}

	return output
}
