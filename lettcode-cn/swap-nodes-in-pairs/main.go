package main

 type ListNode struct {
	     Val int
	     Next *ListNode
	 }
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur:= head.Next, head
	cur.Next = pre
	for pre.Next != nil && pre.Next.Next!=nil{
		f,s := pre.Next, pre.Next.Next
		pre.Next, f.Next, s.Next, pre = s, s.Next, f, f
	}

	return cur
}

func main()  {
	n1 := ListNode{
		Val:  1,
		Next: nil,
	}
	n2 := ListNode{
		Val:  2,
		Next: nil,
	}
	n3 := ListNode{
		Val:  3,
		Next: nil,
	}
	n4 := ListNode{
		Val:  4,
		Next: nil,
	}
	n1.Next = &n2
	n2.Next= &n3
	n3.Next = &n4

	swapPairs(&n1)
}