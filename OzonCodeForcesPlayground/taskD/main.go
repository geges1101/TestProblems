package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func rearrange(table [][]int, column int, n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, len(table[0]))
	}
	m := make(map[int]int)
	order := make([]int, n)
	for i := 0; i != n; i++ {
		order[i] = table[i][column]
		m[table[i][column]] = i
	}
	sort.Ints(order)
	for i := 0; i != n; i++ {
		if order[i] != table[i][column] {
			row := m[order[i]]
			for j := 0; j != len(table[0]); j++ {
				res[i][j] = table[row][j]
			}
		} else {
			for j := 0; j != len(table[0]); j++ {
				res[i][j] = table[i][j]
			}
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tables int
	fmt.Fscan(in, &tables)

	for t := 0; t < tables; t++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		table := make([][]int, n)
		for i := range table {
			table[i] = make([]int, m)
		}

		for i := 0; i != n; i++ {
			for j := 0; j != m; j++ {
				var curr int
				fmt.Fscan(in, &curr)
				table[i][j] = curr
			}
		}

		var clicks, last int
		fmt.Fscan(in, &clicks)

		for i := 0; i != clicks; i++ {
			var click int
			fmt.Fscan(in, &click)
			if last != click {
				table = rearrange(table, click-1, n)
			}
			last = click
		}

		for row := 0; row < n; row++ {
			for column := 0; column < m; column++ {
				fmt.Print(table[row][column], " ")
			}
			fmt.Print("\n")
		}
		fmt.Println()
	}
}

//1
//3 3
//2 11 72
//99 11 13
//2 8 13
//5
//2 3 2 1 2
