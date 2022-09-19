package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumLeftLeaves(root, false)
}
func sumLeftLeaves(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	v := sumLeftLeaves(root.Left, true) + sumLeftLeaves(root.Right, false)
	if isLeft && root.Left == nil && root.Right == nil {
		v += root.Val
	}
	return v
}
