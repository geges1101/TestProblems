package main

import "math"

type MaxStack struct {
	data []int
	max  []int
}

func Constructor() MaxStack {
	return MaxStack{}
}

func (ms *MaxStack) Push(x int) {
	ms.data = append(ms.data, x)
	if len(ms.max) == 0 {
		ms.max = append(ms.max, x)
	} else {
		ms.max = append(ms.max,
			int(math.Max(float64(x),
				float64(ms.max[len(ms.max)-1]))))
	}
}

func (ms *MaxStack) Pop() int {
	ms.max = ms.max[:len(ms.max)-1]
	val := ms.data[len(ms.data)-1]
	ms.data = ms.data[:len(ms.data)-1]
	return val
}

func (ms *MaxStack) Top() int {
	return ms.data[len(ms.data)-1]
}

func (ms *MaxStack) PeekMax() int {
	return ms.max[len(ms.max)-1]
}

func (ms *MaxStack) PopMax() int {
	maxVal := ms.PeekMax()
	var tmp []int

	for ms.Top() != maxVal {
		tmp = append(tmp, ms.Pop())
	}

	ms.Pop()

	for i := len(tmp) - 1; i >= 0; i-- {
		ms.Push(tmp[i])
	}

	return maxVal
}
