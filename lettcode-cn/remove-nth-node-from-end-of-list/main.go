package main

type ListNode struct {
      Val int
      Next *ListNode
  }
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cur := head
    i:=1
    nth := head
	var pre *ListNode
	for cur.Next!=nil {
    	if i == n {
    		break
        }
        i++
        cur = cur.Next
    }
	for cur.Next!=nil {
		nth , pre= nth.Next, nth
		cur = cur.Next
	}
	if nth == head {
		return head.Next
	}
	pre.Next = nth.Next
	return head
}

func main()  {
	n1 := ListNode{Val:1}
	//n2 := ListNode{Val:2}
	//n3 := ListNode{Val:3}
	//n4 := ListNode{Val:4}
	//n5 := ListNode{Val:5}
	//n1.Next = &n2
	//n2.Next = &n3
	//n3.Next = &n4
	//n4.Next = &n5
	removeNthFromEnd(&n1, 1)
}