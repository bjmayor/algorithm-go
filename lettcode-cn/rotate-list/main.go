package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	n := 0
	node := head
	for node != nil {
		n++
		if node.Next == nil {
			node.Next = head
			break
		}
		node = node.Next
	}
	k = k % n
	nh := n - k
	node = head
	for i := 1; i < nh; i++ {
		node = node.Next
	}
	nhead := node.Next
	node.Next = nil
	return nhead
}

func main() {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	rotateRight(n1, 2)
}
