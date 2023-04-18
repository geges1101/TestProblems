package main

func minWindow(s, t string) string {
	rem := 0
	dict := make(map[byte]int)
	for i := range t {
		rem++
		dict[t[i]]++
	}
	if rem > len(s) {
		return ""
	}
	res := string(make([]byte, len(s)))
	start, end := 0, 0
	for end < len(s) {
		if v, ok := dict[s[end]]; ok {
			if v > 0 {
				rem--
			}
			dict[s[end]]--
		}
		for rem <= 0 {
			if len(res) >= len(s[start:end+1]) {
				res = s[start : end+1]
			}
			if _, ok := dict[s[start]]; ok {
				dict[s[start]]++
				if dict[s[start]] > 0 {
					rem++
				}
			}
			start++
		}
		end++
	}
	if res == string(make([]byte, len(s))) {
		return ""
	}
	return res
}
