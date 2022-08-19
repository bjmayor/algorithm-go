package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	s := 0
	q := []*TreeNode{}
	if root != nil {
		q = append(q, root)
	}

	for len(q) > 0 {
		tmp := q
		s = 0
		for _, node := range tmp {
			s += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[len(tmp):]
	}
	return s
}
