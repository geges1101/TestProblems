package main

import (
	"fmt"
)

func missingNum(nums []int) int {
	big, small := 0, 0
	n := len(nums)
	for i := 0; i < n; i++ {
		small += nums[i]
	}
	for i := 0; i <= n; i++ {
		big += i
	}
	return big - small
}

func main() {
	fmt.Print(missingNum([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
