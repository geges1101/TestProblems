package main

import (
	"bufio"
	"fmt"
	"os"
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

	second = begin[0] <= 24 && end[0] <= 24 &&
		begin[1] <= 59 && end[1] <= 59 &&
		begin[2] <= 59 && end[2] <= 59

	return first && second
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
		begin := make([]int, 3)
		end := make([]int, 3)

		flag := 1
		for i := 0; i != n; i++ {
			var line string
			fmt.Fscan(in, &line)
			begin[0], _ = strconv.Atoi(line[0:2])
			begin[1], _ = strconv.Atoi(line[3:5])
			begin[2], _ = strconv.Atoi(line[6:8])
			end[0], _ = strconv.Atoi(line[9:11])
			begin[1], _ = strconv.Atoi(line[12:14])
			begin[2], _ = strconv.Atoi(line[16:18])
			if !validate(begin, end) {
				break
				flag = -1
				fmt.Println("NO")
			}
		}

		if flag != -1 {
			fmt.Println("YES")
		}
	}
}
