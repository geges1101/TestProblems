package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var res []string
	fmt.Fscan(in, &n)

	first := "21122"
	second := "2122"

	for i := 0; i < n; i++ {
		var line string
		fmt.Fscan(in, &line)
		format := 0
		curr := ""
		for j := 0; j < len(line); j++ {
			last := 0
			if format == 0 {
				if unicode.IsNumber(rune(line[j])) {
					curr += "1"
					last = 1
				} else {
					curr += "2"
					last = 2
				}
				if len(curr) == 3 {
					format = last
				}
				if curr != first[j-len(curr):j] || curr != second[j-len(curr):j] {
					res = []string{"-"}
					break
				}
			} else {
				if unicode.IsNumber(rune(line[j])) {
					curr += "1"
				} else {
					curr += "2"
				}
				if format == 1 {
					if curr == first[j-len(curr):j] {
						if len(curr) == 5 {
							res = append(res, line[j-len(curr):j])
							curr = ""
						}
					} else {
						res = []string{"-"}
						break
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, res[i])
	}
}
