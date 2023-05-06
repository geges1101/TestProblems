package main

func removeDuplicates(nums []int) int {
	idx := 0
	for i:= 0; i != len(nums); i++{
		if nums[idx] != nums[i]{
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}
