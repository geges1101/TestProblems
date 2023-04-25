package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var sum = 0

	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	l,r := 0,0

	if low <= root.Val {
		l = rangeSumBST(root.Left, low, high)
	}

	if high >= root.Val {
		r = rangeSumBST(root.Right, low, high)
	}

	return sum + l + r
}
