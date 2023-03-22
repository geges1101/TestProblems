package main

import (
	"math"
	"sort"
)

func maxGuests(guests [][]int) int {
	arriving, leaving := make(map[int]int), make(map[int]int)

	var days []int
	for _, g := range guests {
		arriving[g[0]]++
		leaving[g[1]]++
		days = append(days, g[0])
		days = append(days, g[1])
	}

	sort.Ints(days)
	res := 0
	curr := 0

	for _, day := range days {
		curr -= leaving[day]
		curr += arriving[day]

		res = int(math.Max(float64(res), float64(curr)))
	}

	return res
}
