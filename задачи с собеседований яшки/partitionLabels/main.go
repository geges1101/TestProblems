package main

func partitionLabels(s string) []int {
	res := []int{}
	last := [26]int{}
	for i := range s {
		last[s[i]-'a'] = i
	}
	start, end := 0, 0
	for i := range s {
		end = max(end, last[s[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
