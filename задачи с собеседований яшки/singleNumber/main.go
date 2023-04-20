package main

func singleNumber(nums []int) int {
	mp := make(map[int]int,len(nums)/2)
	res := 0

	for _,v := range nums{
		mp[v]++
		if mp[v] == 1{
			res += v
		} else{
			res -= v
		}
	}
	return res
}
