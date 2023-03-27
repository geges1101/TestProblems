package main

func moveZeroes(nums []int) {
	zeroes, idx, n := 0, 0, len(nums)
	for i := 0; i != n; i++ {
		if nums[i] == 0 {
			zeroes++
		} else {
			nums[idx] = nums[i]
			idx++
		}
	}
	for i := n - zeroes; i != n; i++ {
		nums[i] = 0
	}
}
