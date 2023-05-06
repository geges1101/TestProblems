package sameTree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p,q)
}

func dfs(node1,node2 *TreeNode)bool{
	if node1 == nil && node2 == nil{
		return true
	}
	if node1 == nil || node2 == nil{
		return false
	}
	if node1.Val != node2.Val{
		return false
	}
	left := dfs(node1.Left,node2.Left)
	right := dfs(node1.Right,node2.Right)
	return left && right
}
