package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var l int
var dummy *ListNode
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l = 0
	dummy = head
	traverse(dummy,nil,n)
	return dummy
}

func traverse(node,prev *ListNode, n int){
	if node == nil{
		return
	}
	traverse(node.Next, node, n)
	l++
	if(l == n){
		if prev == nil{
			dummy = node.Next
			return
		}
		prev.Next = node.Next
	}
}
