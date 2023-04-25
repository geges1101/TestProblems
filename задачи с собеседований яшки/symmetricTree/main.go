package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left,root.Right)
}

func dfs(left,right *TreeNode) bool{
	if left == nil && right == nil{
		return true
	}
	if left == nil || right == nil{
		return false
	}
	if left.Val != right.Val{
		return false
	}
	return dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
}
