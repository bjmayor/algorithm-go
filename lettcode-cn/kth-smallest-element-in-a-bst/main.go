package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	values := walkTree(root)
	return values[k-1]
}

func walkTree(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	l := walkTree(node.Left)
	r := walkTree(node.Right)
	l = append(l, node.Val)
	l = append(l, r...)
	return l
}
