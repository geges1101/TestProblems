package main

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n, n)
	l, r := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if nums[l]*nums[l] > nums[r]*nums[r] {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}

	return res
}
