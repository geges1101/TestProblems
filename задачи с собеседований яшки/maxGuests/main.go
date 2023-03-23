package main

import (
	"math"
	"sort"
)

//Даны даты заезда и отъезда каждого гостя. Для каждого гостя дата заезда строго раньше даты отъезда (то есть каждый гость останавливается хотя бы на одну ночь). В пределах одного дня считается, что сначала старые гости выезжают, а затем въезжают новые. Найти максимальное число постояльцев, которые одновременно проживали в гостинице (считаем, что измерение количества постояльцев происходит в конце дня).
//
//sample = [ (1, 2), (1, 3), (2, 4), (2, 3), ]

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
