package main

import "math"

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0

	for _, num := range nums {
		sum += num

		max = int(math.Max(float64(max), float64(sum)))

		if sum < 0 {
			sum = 0
		}
	}
	return max
}
