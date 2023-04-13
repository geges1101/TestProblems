package main

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	water := 0
	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(height[left], maxLeft)
			water += maxLeft - height[left]
		} else {
			right--
			maxRight = max(height[right], maxRight)
			water += maxRight - height[right]
		}
	}
	return water
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
