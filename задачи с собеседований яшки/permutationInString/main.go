package main

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var dict, target [26]int
	i := 0

	for ; i < n; i++ {
		target[s1[i]-'a']++
		dict[s2[i]-'a']++
	}

	if dict == target {
		return true
	}

	for ; i < m; i++ {
		dict[s2[i]-'a']++
		dict[s2[i-n]-'a']--
		if dict == target {
			return true
		}
	}

	return false
}
