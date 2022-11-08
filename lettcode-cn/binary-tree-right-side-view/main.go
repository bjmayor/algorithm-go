package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		right := q[len(q)-1]
		res = append(res, right.Val)
		var tmp []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		q = tmp
	}
	return res
}

func main() {
	//tree := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	//fmt.Println(rightSideView(tree))
	fmt.Println(rightSideView(nil))
}
