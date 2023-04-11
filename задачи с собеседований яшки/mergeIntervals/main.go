package main

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	var res [][]int
	begin, end := intervals[0][0], intervals[0][1]
	for i := 1; i != len(intervals); i++ {
		if intervals[i][0] <= end {
			end = int(math.Max(float64(end), float64(intervals[i][1])))
		} else {
			res = append(res, []int{begin, end})
			begin, end = intervals[i][0], intervals[i][1]
		}
	}
	res = append(res, []int{begin, end})
	return res
}
