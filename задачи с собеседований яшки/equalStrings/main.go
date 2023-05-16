package main

import (
	"fmt"
	"math"
)

//Даны две строки.
//
//Написать функцию, которая вернёт True, если из первой строки можно получить вторую, совершив не более 1 изменения (== удаление / замена символа).

func equalStrings(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	diff := math.Abs(float64(n - m))
	if diff > 1 {
		return false
	}
	if m > n {
		n, m = m, n
		s1, s2 = s2, s1
	}

	i, j, miss := 0, 0, 0
	second := true
	for j < m {
		if s1[i] != s2[j] {
			if miss > 0 {
				return false
			}
			miss++
			x, y := i+1, j
			for y < m && x < n {
				if s1[x] != s2[y] {
					second = false
				}
				x++
				y++
			}
			break
		}
		i++
		j++
	}
	return second || (miss < 1 || diff < 1)
}

func main() {
	fmt.Println(equalStrings("aa", ""))
}
