package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	heights := strings.Split(input, " ")
	n := len(heights)

	if n <= 1 {
		fmt.Println("YES")
		return
	}

	if heights[0] <= heights[n-1] {
		for i := 1; i < len(heights); i++ {
			if heights[i] < heights[i-1] {
				fmt.Print("NO")
				return
			}
		}
	} else {
		for i := 1; i < len(heights); i++ {
			if heights[i] > heights[i-1] {
				fmt.Print("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
