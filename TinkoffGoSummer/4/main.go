package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n+1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 1; i <= n; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &a[i])
	}

	cnt := make(map[int]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	ans := -1

	for l := 1; l <= n; l++ {
		prevCnt := cnt[a[l]]
		for i := 0; i < len(pq); i++ {
			if pq[i].value == a[l] && pq[i].priority == prevCnt {
				heap.Remove(&pq, i)
				break
			}
		}
		cnt[a[l]]++
		heap.Push(&pq, &Item{value: a[l], priority: cnt[a[l]]})

		minCnt := pq[0].priority
		maxCnt := pq[len(pq)-1].priority

		if minCnt == maxCnt {
			ans = l
			continue
		}

		if len(pq) == 2 && maxCnt-minCnt == 1 && maxCnt != n {
			ans = l
			continue
		}

		if len(pq) == 1 && minCnt == 1 {
			ans = l
			continue
		}

		if len(pq) == 2 && minCnt == 1 {
			ans = l
			continue
		}

		if len(pq) <= 2 && minCnt == 1 {
			ans = l
			continue
		}
		
		if len(pq) <= 2 && maxCnt == n-1 {
			ans = l
			continue
		}
	}

	fmt.Println(ans)
}
