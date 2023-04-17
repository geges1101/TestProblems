package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, v := range nums {
		mp[v] = i
	}
	for first, v := range nums {
		curr := target - v
		second, ok := mp[curr]
		if ok && first != second {
			return []int{first, second}
		}
	}
	return []int{-1, -1}
}
