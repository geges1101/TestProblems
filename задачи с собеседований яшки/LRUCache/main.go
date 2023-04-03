package main

type LRUCache struct {
	cache map[int]*Node
	cap   int
	head  *Node
	tail  *Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*Node)
	head := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, nil, nil}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		cache,
		capacity,
		head,
		tail}
}

func (this *LRUCache) RemNode(node *Node) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) AddNode(node *Node) {
	node.Prev = this.head
	node.Next = this.head.Next
	node.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.RemNode(node)
		this.AddNode(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.Val = value
		this.RemNode(node)
		this.AddNode(node)
	} else {
		if this.cap == 0 {
			least := this.tail.Prev
			this.RemNode(least)
			delete(this.cache, least.Key)
		} else {
			this.cap--
		}

		node := &Node{key, value, nil, nil}
		this.AddNode(node)
		this.cache[key] = node
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
