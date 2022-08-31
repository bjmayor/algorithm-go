package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	m := nums[0]
	idx := 0
	for i, v := range nums[1:] {
		if v > m {
			m = v
			idx = i + 1
		}
	}
	n := &TreeNode{
		Val: m,
	}
	n.Left = constructMaximumBinaryTree(nums[0:idx])
	n.Right = constructMaximumBinaryTree(nums[idx+1:])
	return n
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}
