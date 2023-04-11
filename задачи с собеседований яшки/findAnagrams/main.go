package main

func findAnagrams(s string, p string) (res []int) {
	lnP := len(p)
	if len(s) < lnP {
		return res
	}

	const (
		abcLen = 26
		abcBeg = 'a'
	)

	sCnt, pCnt := [abcLen]int{}, [abcLen]int{}
	for _, rn := range p {
		pCnt[rn-abcBeg]++
	}

	for i := range s {
		sCnt[s[i]-abcBeg]++
		if i >= lnP {
			sCnt[s[i-lnP]-abcBeg]--
		}
		if sCnt == pCnt {
			res = append(res, i-lnP+1)
		}
	}

	return res
}
