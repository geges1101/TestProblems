package main

import "fmt"

func FindMaxConsecutiveOnes(nums []int) int {
	// write your code here
	start, i, zeroes, res := 0, 0, 0, 0
	for i < len(nums) {
		if nums[i] == 0 {
			zeroes++
		}
		for zeroes > 1 {
			if nums[start] == 0 {
				zeroes--
			}
			start++
		}
		res = max(res, i-start+1)
		i++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Print(FindMaxConsecutiveOnes([]int{1, 1}))
}
