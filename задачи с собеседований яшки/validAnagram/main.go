package validAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	dict := make(map[byte]int)
	for i := range s{
		dict[s[i]]++
		dict[t[i]]--
	}
	for _, v := range dict{
		if v != 0{
			return false
		}
	}
	return true
}
