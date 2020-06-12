package main
 type ListNode struct {
	     Val int
	     Next *ListNode
	 }

//迭代解法
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	pre, cur := nil, head
	for cur != nil {
		ppre := pre
		pre.Next, pre, cur = pre, cur, cur.Next
		pre.Next = ppre
	}
	return pre
}

//递归解法
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	next := reverseList2(head.Next)
	head.Next.Next = head
	head.Next= nil
	return next
}

