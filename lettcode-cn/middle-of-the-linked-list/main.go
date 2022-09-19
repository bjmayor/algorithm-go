package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	p := head
	pp := head
	for pp.Next != nil {
		p = p.Next
		pp = pp.Next
		if pp.Next != nil {
			pp = pp.Next
		} else {
			break
		}
	}
	return p
}
