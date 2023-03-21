package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m, k float64
	fmt.Scan(&n, &m, &k)
	fmt.Println(math.Round(n * k / m))
}
