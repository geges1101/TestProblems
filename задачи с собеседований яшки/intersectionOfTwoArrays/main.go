package main

func intersect(nums1 []int, nums2 []int) []int {
	mapping := make(map[int]int)
	var intersection []int
	for _, v := range nums1 {
		mapping[v]++
	}

	for _, v := range nums2 {
		if mapping[v] > 0 {
			intersection = append(intersection, v)
			mapping[v]--
		}
	}
	return intersection
}
