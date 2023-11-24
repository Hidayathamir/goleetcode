package main

func productExceptSelf(nums []int) []int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	output := make([]int, len(nums))
	// iterasi 1, prefix
	x := 1
	var y, prevNumber int
	for i := 1; i < len(nums)-1; i++ {
		prevNumber = nums[i-1]
		y = prevNumber * x
		output[i] = y
		x = y
	}

	// iterasi 2, postfix
	x = 1
	var nextNumber int
	for i := len(nums) - 2; i > 0; i-- {
		nextNumber = nums[i+1]
		y = nextNumber * x
		output[i] = y * output[i]
		x = y
	}

	return output[1 : len(output)-1]
}
