package main

import (
	"fmt"
	"sort"
)

//Слияние отрезков:
//
//Вход: [1, 3] [100, 200] [2, 4]
//Выход: [1, 4] [100, 200]

func merge(rngs [][]int) [][]int {
	sort.Slice(rngs, func(i, j int) bool {
		return rngs[i][0] < rngs[j][0]
	})

	var res [][]int
	start, end := rngs[0][0], rngs[0][1]
	for _, rng := range rngs {
		if rng[0] < end {
			end = max(end, rng[1])
		} else {
			res = append(res, []int{start, end})
			start, end = rng[0], rng[1]
		}
	}
	res = append(res, []int{start, end})
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
	fmt.Print(merge([][]int{{1, 3}, {100, 200}, {2, 4}}))
}
