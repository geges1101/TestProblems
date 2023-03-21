package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		var curr int64
		fmt.Scan(&curr)
		nums[i] = curr
	}
	pref := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		pref[i] = pref[i-1] + nums[i]
	}
	back := make(map[int64]int)
	back[0] = 0
	last := 0
	var result int64 = 0
	for i := 1; i <= n; i++ {
		_, ok := back[pref[i]]
		if ok {
			idx := back[pref[i]] + 1
			l, r := idx-last, n-i+1
			result += int64(l * r)
			back[pref[i]] = i
			last = idx
		} else {
			back[pref[i]] = i
		}
	}
	fmt.Println(result)
}
