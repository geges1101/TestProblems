package main

import (
	"fmt"
	"sort"
	"strconv"
)

func segment(nums []int) string {
	sort.Ints(nums)
	var result string
	start := nums[0]
	end := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			if start == end {
				result += strconv.Itoa(start) + ","
			} else {
				result += strconv.Itoa(start) + "-" + strconv.Itoa(end) + ","
			}
			start = nums[i]
			end = nums[i]
		}
	}
	if start == end {
		result += strconv.Itoa(start)
	} else {
		result += strconv.Itoa(start) + "-" + strconv.Itoa(end)
	}

	return result
}

func main() {
	fmt.Print(segment([]int{1, 4, 5, 2, 3, 9, 8, 11, 0}))
}
