package main

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var q []*TreeNode
	if root == nil {
		return nil
	}
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		var total float64
		for i:=0;i<size;i++ {
			first := q[0]
			q = q[1:]
			total += float64(first.Val)
			if first.Left != nil {
				q = append(q, first.Left)
			}
			if first.Right != nil {
				q = append(q, first.Right)
			}
		}
		res = append(res, total/float64(size))
	}
	return res
}
