package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var a []*TreeNode
	a = append(a, root)
	level := 0
	for len(a) > 0 {
		level+=1
		size := len(a)
		for i:=0;i<size;i++ {
			node := a[0]
			a = a[1:]
			if node.Right==nil && node.Left == nil {
				return level
			}
			if node.Left != nil {
				a = append(a, node.Left)
			}

			if node.Right != nil {
				a = append(a, node.Right)
			}
		}
	}
	return level
}

