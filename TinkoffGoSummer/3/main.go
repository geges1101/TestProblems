package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanln(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	l, r, count := 0, 0, 0
	letters := make(map[string]int)
	min := n
	for r < n {
		if letters["a"] > 0 &&
			letters["b"] > 0 &&
			letters["c"] > 0 &&
			letters["d"] > 0 {
			min = Min(r-l, min)
			letters[string(s[l])]--
			l++
			count++
		} else {
			letters[string(s[r])]++
			if letters["a"] > 0 &&
				letters["b"] > 0 &&
				letters["c"] > 0 &&
				letters["d"] > 0 {
				min = Min(r-l+1, min)
				count++
			}
			r++
		}
	}
	if count > 0 {
		fmt.Println(min)
	} else {
		fmt.Print(-1)
	}
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
