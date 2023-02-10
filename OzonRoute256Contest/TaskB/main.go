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
	var res []string
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		ships := []int{
			1: 4,
			2: 3,
			3: 2,
			4: 1,
		}
		for j := 0; j < 10; j++ {
			var curr int
			fmt.Fscan(in, &curr)
			ships[curr]--
		}
		f := true
		for _, val := range ships {
			if val != 0 {
				res = append(res, "NO")
				f = false
				break
			}
		}
		if f {
			res = append(res, "YES")
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, res[i])
	}
}