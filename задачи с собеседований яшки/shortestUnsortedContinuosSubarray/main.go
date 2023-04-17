package main

import "math"

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1

	for low < len(nums)-1 && nums[low] <= nums[low+1] {
		low++
	}
	if low == len(nums)-1 {
		return 0
	}

	for high > 0 && nums[high] >= nums[high-1] {
		high--
	}

	subMin, subMax := math.MaxInt, math.MinInt

	for i := low; i <= high; i++ {
		if nums[i] < subMin {
			subMin = nums[i]
		}
		if nums[i] > subMax {
			subMax = nums[i]
		}
	}

	for low > 0 && nums[low-1] > subMin {
		low--
	}
	for high < len(nums)-1 && nums[high+1] < subMax {
		high++
	}

	return high - low + 1
}
