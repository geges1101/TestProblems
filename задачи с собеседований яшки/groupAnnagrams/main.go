package main

import (
	"fmt"
)

//Sample Input ["eat", "tea", "tan", "ate", "nat", "bat"]
//Sample Output [ ["ate", "eat", "tea"], ["nat", "tan"], ["bat"] ]
//
//Т.е. сгруппировать слова по "общим буквам".

func organize(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}
	var res [][]string
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Print(organize([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
