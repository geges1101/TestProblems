package main

import (
	"fmt"
	"math"
)

func getSquares(numbers []int) []int {
	n := len(numbers)
	left := 0
	right := n - 1
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		if math.Abs(float64(numbers[left])) >= math.Abs(float64(numbers[right])) {
			result[i] = numbers[left] * numbers[left]
			left++
		} else {
			result[i] = numbers[right] * numbers[right]
			right--
		}
	}
	return result
}

func main() {
	fmt.Print(getSquares([]int{-12, -7, -5, -1, 1, 2, 8, 34}))
}
