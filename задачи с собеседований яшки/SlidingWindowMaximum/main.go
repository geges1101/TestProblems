package main

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var dq []int
	l, r := 0, 0

	for r < len(nums) {
		for len(dq) > 0 && nums[dq[len(dq)-1]] < nums[r] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, r)

		if l > dq[0] {
			dq = dq[1:]
		}

		if (r + 1) >= k {
			result = append(result, nums[dq[0]])
			l += 1
		}
		r += 1
	}
	return result
}
