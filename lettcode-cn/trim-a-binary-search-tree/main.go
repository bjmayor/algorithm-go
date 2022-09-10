package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low || root.Val > high {
		left := trimBST(root.Left, low, high)
		if left != nil {
			return left
		} else {
			return trimBST(root.Right, low, high)
		}
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func main() {
	n2 := &TreeNode{
		Val: 2,
	}
	n1 := &TreeNode{
		Val: 0,
	}
	n0 := &TreeNode{
		Val:   1,
		Left:  n1,
		Right: n2,
	}
	trimBST(n0, 1, 2)
}
