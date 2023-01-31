package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func validate(begin []int, end []int) bool {
	first := false
	second := false
	if begin[0] == end[0] {
		if begin[1] == end[1] {
			if begin[2] == end[2] {
				first = true
			} else {
				first = begin[2] < end[2]
			}
		} else {
			first = begin[1] < end[1]
		}
	} else {
		first = begin[0] < end[0]
	}

	second = begin[0] <= 23 && end[0] <= 23 &&
		begin[1] <= 59 && end[1] <= 59 &&
		begin[2] <= 59 && end[2] <= 59

	return first && second
}

type Pair struct {
	begin []int
	end   []int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var sets int
	fmt.Fscan(in, &sets)

	for s := 0; s < sets; s++ {
		var n int
		fmt.Fscan(in, &n)
		all := make([]Pair, n)

		flag := 1
		for i := 0; i != n; i++ {
			begin := make([]int, 3)
			end := make([]int, 3)
			var line string
			fmt.Fscan(in, &line)
			begin[0], _ = strconv.Atoi(line[0:2])
			begin[1], _ = strconv.Atoi(line[3:5])
			begin[2], _ = strconv.Atoi(line[6:8])
			end[0], _ = strconv.Atoi(line[9:11])
			end[1], _ = strconv.Atoi(line[12:14])
			end[2], _ = strconv.Atoi(line[15:17])
			all[i] = Pair{begin: begin, end: end}
		}

		sort.Slice(all, func(i, j int) bool {
			return all[i].begin[0]*60*60+all[i].begin[1]*60+all[i].begin[2] <
				all[j].begin[0]*60*60+all[j].begin[1]*60+all[j].begin[2]
		})

		for _, v := range all {
			if !validate(v.begin, v.end) {
				flag = -1
				fmt.Println("NO")
				break
			}
		}

		for i := 0; i < n-1; i++ {
			if all[i].end[0]*60*60+all[i].end[1]*60+all[i].end[2] >=
				all[i+1].begin[0]*60*60+all[i+1].begin[1]*60+all[i+1].begin[2] {
				flag = -1
				fmt.Println("NO")
				break
			}
		}
		if flag != -1 {
			fmt.Println("YES")
		}
	}
}
