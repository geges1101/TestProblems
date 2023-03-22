package main

import "fmt"

//Дан массив из нулей и единиц. Нужно определить, какой максимальный по длине подинтервал единиц можно получить, удалив ровно один элемент массива.
//
//[1, 1, 0]

func maxOnes(nums []int) int {
	res, start, end, n := 0, 0, 0, len(nums)

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			start = i
			end = start
			for end < n {
				if nums[end] == 1 {
					end++
				} else {
					break
				}
			}
			i = end
			res = max(res, end-start)
		}
	}

	for i := 1; i < n-1; i++ {
		if nums[i-1] == 1 && nums[i] == 0 && nums[i+1] == 1 {
			l, r := i-1, i+1
			countl, countr := 0, 0
			for l >= 0 {
				if nums[l] == 1 {
					countl++
					l--
				} else {
					break
				}
			}
			for r < n {
				if nums[r] == 1 {
					countr++
					r++
				} else {
					break
				}
			}
			res = max(res, countl+1+countr)
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Print(maxOnes([]int{1, 0, 1, 0, 1, 1, 1, 0, 1}))
}
