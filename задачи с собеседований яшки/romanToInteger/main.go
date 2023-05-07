package romanToInteger

func romanToInt(s string) int {
	n := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var res int

	for j, i := range s {
		if j+1 < len(s) && n[rune(s[j+1])] > n[i] {
			res -= n[i]
		} else {
			res += n[i]
		}
	}

	return res
}
