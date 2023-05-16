package main

import (
	"fmt"
	"strconv"
)

//Дан массив точек с целочисленными координатами (x, y).
//Определить, существует ли вертикальная прямая, делящая точки на 2 симметричных относительно этой прямой множества.
//Note: Для удобства точку можно представлять не как массив [x, y], а как объект {x, y}

type Point struct {
	x, y int
}

func isSymmetrical(points []Point) bool {
	max, min := points[0].x, points[0].x
	mp := make(map[string]bool)
	for _, p := range points {
		if p.x > max {
			max = p.x
		}
		if p.x < min {
			min = p.x
		}
		mp[strconv.Itoa(p.x)+"."+strconv.Itoa(p.y)] = true
	}
	center := min + max
	for _, p := range points {
		if mp[strconv.Itoa(center-p.x)+"."+strconv.Itoa(p.y)] != true {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(isSymmetrical([]Point{{
		x: 0,
		y: 3,
	}, {
		x: 6,
		y: 3,
	}, {
		x: 3,
		y: 1,
	}}))
}
