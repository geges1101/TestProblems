package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var length int
		fmt.Fscan(in, &length)
		visited := make(map[int]int)
		var res [][]int
		workers := make([]int, length)
		for j := 0; j != length; j++ {
			var curr int
			fmt.Fscan(in, &curr)
			workers[j] = curr
		}
		for j := 0; j != length; j++ {
			if visited[j] == 0 {
				skill := workers[j]
				min := math.MaxFloat64
				visited[j] = 1
				var pair, idx int
				for v := 0; v < length; v++ {
					if visited[v] == 0 && v != j {
						pair = workers[v]
						curr := math.Abs(float64(skill - pair))
						if min > curr {
							min = curr
							idx = v
						}
					}
				}
				visited[idx] = 1
				res = append(res, []int{j + 1, idx + 1})
			}
		}
		fmt.Fprintln(out, res)
	}
}
