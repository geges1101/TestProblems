package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var res [][]int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)
		nums := make([]int, m)
		data := make(map[int]int)
		for j := 0; j < m; j++ {
			var curr int
			fmt.Fscan(in, &curr)
			nums[j] = curr
		}
		sorted := nums
		sort.Ints(sorted)
		places := make([]int, m)
		place := 1
		streak := 0
		last := sorted[0]
		for j := 0; j < m; j++ {
			if sorted[j]-last > 1 {
				place += streak
				streak = 0
			}
			streak++
			places[j] = place
			data[sorted[j]] = places[j]
		}
		for j := 0; j < m; j++ {
			nums[j] = data[nums[j]]
		}
		res = append(res, nums)
	}

	for row := 0; row < n; row++ {
		l := len(res[row])
		for column := 0; column < l; column++ {
			fmt.Print(res[row][column], " ")
		}
		fmt.Print("\n")
	}
}
