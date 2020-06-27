package main

  type ListNode struct {
      Val int
      Next *ListNode
  }
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := l1
	h2 := l2

	var head *ListNode
	var cur *ListNode
	for h1!=nil && h2!=nil {
		s := h1
		if h2.Val < h1.Val {
			s = h2
			h2 = h2.Next
		} else {
			h1 = h1.Next
		}
		if head == nil {
			head = s
			cur = s
		}
		cur.Next, cur = s, s
	}
	if h1!=nil {
		if cur != nil {
			cur.Next = h1
		} else {
			head = h1
		}
	}
	if h2!=nil {
		if cur != nil {
			cur.Next = h2
		} else {
			head = h2
		}
	}
	return head
}
