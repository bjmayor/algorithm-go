package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	} else {
		if root.Val < val {
			n := &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			n.Left = root
			return n
		} else {
			root.Right = insertIntoMaxTree(root.Right, val)
		}
	}

	return root
}
