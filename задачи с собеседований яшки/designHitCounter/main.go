package main

type HitCounter struct {
	ts   [300]int
	hits [300]int
}

func Constructor() *HitCounter {
	return &HitCounter{}
}

func (hc *HitCounter) hit(timestamp int) {
	i := timestamp % 300
	if hc.ts[i] != timestamp {
		hc.ts[i] = timestamp
		hc.hits[i] = 1
	} else {
		hc.hits[i]++
	}
}

func (hc *HitCounter) getHits(timestamp int) int {
	result := 0
	for i := 0; i < len(hc.hits); i++ {
		if timestamp-hc.ts[i] < 300 {
			result += hc.hits[i]
		}
	}
	return result
}
