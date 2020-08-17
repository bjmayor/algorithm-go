package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if ok, l := height(root.Left); ok {
		if ok, r := height(root.Right); ok {
			return l-r <= 1 && l-r >= -1
		}
	}
	return false
}
func height(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lok, l := height(root.Left);
	rok, r := height(root.Right);
	h := max(l, r)
	h+=1
	if lok && rok {
		if l-r <= 1 && l-r >= -1 {
			return true, h
		}
	}
	return false, h
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
