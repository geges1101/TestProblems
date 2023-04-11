package main

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		begin, end := max(firstList[i][0], secondList[i][0]), min(firstList[i][1], secondList[i][1])
		if begin <= end {
			res = append(res, []int{begin, end})
		}
		if firstList[i][1] < secondList[i][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
