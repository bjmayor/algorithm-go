package main


  type ListNode struct {
      Val int
      Next *ListNode
  }

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    slow := head
    fast := head.Next
    for fast != nil {
       fast = fast.Next
        slow = slow.Next
       if fast != nil {
           fast = fast.Next
       } else {
           break
       }
       if fast == slow {
           return true
       }
    }
    return false
}