package main

import (
	"fmt"
	"math"
)

//Дан массив из нулей и единиц. Нужно определить, какой максимальный по длине подинтервал единиц можно получить, удалив ровно один элемент массива.
//
//[1, 1, 0]

func longestSubarray(nums []int) int {
	var start, curr, max int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			curr++
		}
		for curr > 1 {
			if nums[start] == 0 {
				curr--
			}
			start++
		}
		max = int(math.Max(float64(max), float64(i-start)))
	}
	return max
}

func main() {
	fmt.Print(longestSubarray([]int{1, 0, 1, 0, 1, 1, 1, 0, 1}))
}
