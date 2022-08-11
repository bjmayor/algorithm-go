package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	levels := []*TreeNode{root}
	level := 1
	sum := math.MinInt64
	ret :=1
	for len(levels) > 0 {
		size := len(levels)
		lsum :=0
		for i:=0;i<size;i++ {
			n := levels[0];
			levels = levels[1:]
			if n.Left != nil {
				levels = append(levels, n.Left)
			}
			if n.Right != nil {
				levels = append(levels, n.Right)
			}
			lsum += n.Val
		}
		if lsum > sum {
			ret = level
			sum = lsum
		}
		level++

	}
	return ret
}

func main()  {
	n5 := &TreeNode{
		Val:   -8,
		Left:  nil,
		Right: nil,
	}
	n4 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	n3:= &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	n2:= &TreeNode{
		Val:   7,
		Left:  n4,
		Right: n5,
	}
	root := &TreeNode{
		Val:   1,
		Left:  n2,
		Right: n3,
	}
	println(maxLevelSum(root)) //2
}