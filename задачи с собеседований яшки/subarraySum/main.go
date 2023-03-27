package main

func subarraySum(nums []int, k int) int {
	sum, count := 0, 0
	prevSum := map[int]int{0: 1}

	for _, num := range nums {
		sum += num
		if val, ok := prevSum[sum-k]; ok {
			count += val
		}
		prevSum[sum] += 1
	}
	return count
}
