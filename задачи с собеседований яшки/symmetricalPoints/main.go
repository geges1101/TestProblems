package main

import "fmt"

//Дан массив точек с целочисленными координатами (x, y).
//Определить, существует ли вертикальная прямая, делящая точки на 2 симметричных относительно этой прямой множества.
//Note: Для удобства точку можно представлять не как массив [x, y], а как объект {x, y}

type Point struct {
	x, y int
}

func isSymmetrical(points []Point) bool {
	var max, min float64
	max, min = float64(points[0].x), float64(points[0].x)
	for _, p := range points {
		if float64(p.x) > max {
			max = float64(p.x)
		}
		if float64(p.x) < min {
			min = float64(p.x)
		}
	}
	var centerX = min + (max-min)/2
	l, r := 0, len(points)-1
	for l <= r {
		var leftPoint, rightPoint float64
		leftPoint, rightPoint = float64(points[l].x), float64(points[r].x)
		if leftPoint+(rightPoint-leftPoint)/2 != centerX {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	fmt.Print(isSymmetrical([]Point{{
		x: 0,
		y: 4,
	}, {
		x: 6,
		y: 1,
	}, {
		x: -1,
		y: 1,
	}, {
		x: 5,
		y: 0,
	}}))
}
