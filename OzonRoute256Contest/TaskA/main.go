package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintln(out, "I am sure that I will fill out the form by 23:59 on February 5, 2023.")
}