package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m := make(map[string]int)
	result := []*TreeNode{}
	var inOrder func(root *TreeNode, m map[string]int) string
	inOrder = func(root *TreeNode, m map[string]int) string {
		if root == nil {
			return "#"
		}
		s := fmt.Sprintf("%d-%s-%s", root.Val, inOrder(root.Left, m), inOrder(root.Right, m))
		if m[s] == 1 {
			result = append(result, root)
		}
		m[s]++
		return s
	}
	inOrder(root, m)
	return result
}

func main() {
	n1 := &TreeNode{
		Val: 1,
	}
	n2 := &TreeNode{
		Val: 1,
	}
	n := &TreeNode{
		Val:   2,
		Left:  n1,
		Right: n2,
	}
	fmt.Println(findDuplicateSubtrees(n))
}
