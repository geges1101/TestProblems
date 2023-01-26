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

	var workers int
	fmt.Fscan(in, &workers)

	for w := 0; w < workers; w++ {
		var n int
		fmt.Fscan(in, &n)
		visited := make(map[int]int)
		timeline := make([]int, n)
		last := -1

		for i := 0; i != n; i++ {
			var curr int
			fmt.Fscan(in, &curr)
			timeline[i] = curr
		}

		for _, v := range timeline {
			if visited[v] > 0 && last != v {
				fmt.Println("NO")
				last = -1
				break
			} else {
				visited[v]++
			}
			last = v
		}

		if last != -1 {
			fmt.Println("YES")
		}
	}
}
