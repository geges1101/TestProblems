package main

import "math/rand"

type RandomizedSet struct {
	dict map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		dict: make(map[int]bool),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; !ok {
		this.dict[val] = true
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dict[val]; ok {
		delete(this.dict, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	n := len(this.dict)
	randNumber := rand.Intn(n)
	count := 0

	for k := range this.dict {
		if count == randNumber {
			return k
		}
		count++
	}
	return -1
}
