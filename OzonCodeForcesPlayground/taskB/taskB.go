package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var m, overall int
		fmt.Fscan(in, &m)
		groceries := make(map[int]int)
		for j := 0; j < m; j++ {
			var curr int
			fmt.Fscan(in, &curr)
			if groceries[curr] == 2 {
				groceries[curr] = 0
			} else {
				groceries[curr]++
				overall += curr
			}
		}
		fmt.Fprintln(out, overall)
	}
}
