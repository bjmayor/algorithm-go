package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findQorQ(root,p,q)
}

func findQorQ(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := findQorQ(root.Left,p,q)
	right := findQorQ(root.Right,p,q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}