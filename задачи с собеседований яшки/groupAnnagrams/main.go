package main

import (
	"fmt"
	"sort"
	"strings"
)

//Sample Input ["eat", "tea", "tan", "ate", "nat", "bat"]
//Sample Output [ ["ate", "eat", "tea"], ["nat", "tan"], ["bat"] ]
//
//Т.е. сгруппировать слова по "общим буквам".

func organize(words []string) [][]string {
	dict := make(map[string][]string)
	for _, word := range words {
		s := strings.Split(word, "")
		sort.Strings(s)
		w := strings.Join(s, "")
		dict[w] = append(dict[w], word)
	}
	var res [][]string
	for _, group := range dict {
		res = append(res, group)
	}
	return res
}

func main() {
	fmt.Print(organize([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
