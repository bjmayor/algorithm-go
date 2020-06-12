package main


type ListNode struct {
	Val int
	Next *ListNode
}

/**
假设慢指针走了s步，那么快指针走了2s步。 f= 2s
第一次相遇时，快指针比慢指针多走了n圈，假设一圈的长度是b, 即f = s+nb
两式相减可得s = nb, 而走到入口点时步数一定是a+nb, 那么s再走a步即可到达入口点，从头出发也是走a步到入口点。
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head
	found := false
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next
		} else {
			break
		}
		if fast == slow {
			found = true
			break
		}
	}
	if !found  {
		return nil
	}
	var result = head
	for result != fast {
		result = result.Next
		fast = fast.Next
	}
	return result
}

func main()  {
	n1 := ListNode{
		Val:  1,
		Next: nil,
	}
	//n2 := ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//n3 := ListNode{
	//	Val:  0,
	//	Next: nil,
	//}
	//n4 := ListNode{
	//	Val:  -4,
	//	Next: nil,
	//}
	//n1.Next = &n2
	//n2.Next = &n3
	//n3.Next = &n4
	//n4.Next = &n2
	detectCycle(&n1)
}