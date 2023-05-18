package main

func findAnagrams(s string, p string) (res []int) {
	m := len(p)
	if len(s) < m {
		return res
	}

	dict, target := [26]int{}, [26]int{}
	for _, v := range p {
		target[v-'a']++
	}

	for i := range s {
		dict[s[i]-'a']++
		if i >= m {
			dict[s[i-m]-'a']--
		}
		if dict == target {
			res = append(res, i-m+1)
		}
	}

	return res
}
