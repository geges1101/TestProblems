package sumOfLeftLeaves

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var count int
func sumOfLeftLeaves(root *TreeNode) int {
	count = 0
	dfs(root.Left, 0)
	dfs(root.Right,1)
	return count
}

func dfs(node *TreeNode, prev int){
	if node != nil{
		if node.Left == nil && node.Right == nil && prev == 0{
			count += node.Val
			return
		}
		if node.Left != nil{
			dfs(node.Left,0)
		}
		if node.Right != nil{
			dfs(node.Right, 1)
		}
	}
}
