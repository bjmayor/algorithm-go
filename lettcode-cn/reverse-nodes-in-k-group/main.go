package main


  type ListNode struct {
      Val int
      Next *ListNode
  }

func reverseKGroup(head *ListNode, k int) *ListNode {

    result := head
    tail := head
    var pre *ListNode
    if k<=1 {
        return head
    }
    for tail != nil{
        size :=  1
        for i:=2;i<=k;i++ {
            if tail.Next != nil {
                tail = tail.Next
                size++
            } else if i<k {
                break
            }
        }
        if size == k {
            nextHead := tail.Next
            reverse(head,k)
            head.Next = tail
            if result == head {
                result = tail
                head.Next = nil
                pre = head
            } else {
                pre.Next = tail
                head.Next = nil
                pre = head
            }
            tail = nextHead
            head = nextHead
        } else {
        	pre.Next = head
        	break
        }
    }

    return result
}

func reverse(head *ListNode,k int) *ListNode  {
	var pre *ListNode
   pre, cur:= nil, head
   i := 1
   for cur!= nil && i<=k{
       cur.Next, pre, cur= pre, cur, cur.Next
      i++
   }
   return pre
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
    //n3 := ListNode{
    //	Val:  3,
    //	Next: nil,
    //}
    //n4 := ListNode{
    //	Val:  4,
    //	Next: nil,
    //}
    //n5 := ListNode{
    //    Val:  5,
    //    Next: nil,
    //}
    n1.Next = &n2
    //n2.Next = &n3
    //n3.Next = &n4
    //n4.Next = &n5
    reverseKGroup(&n1, 2)
}