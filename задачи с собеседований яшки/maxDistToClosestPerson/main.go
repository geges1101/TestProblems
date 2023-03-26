package main

func maxDistToClosest(seats []int) int {
	best, zeroes := 0, 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			zeroes++

			if zeroes == i+1 {
				best = max(best, zeroes)
			} else {
				best = max(best, (zeroes+1)/2)
			}
		} else {
			zeroes = 0
		}
	}

	return max(best, zeroes)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
