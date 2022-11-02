package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	data []int
	idx  int
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	right := inOrder(root.Right)
	return append(append(inOrder(root.Left), root.Val), right...)
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		data: inOrder(root),
		idx:  -1,
	}
}

func (this *BSTIterator) Next() int {
	this.idx += 1
	return this.data[this.idx]
}

func (this *BSTIterator) HasNext() bool {
	if this.idx < len(this.data)-1 {
		return true
	}
	return false
}
