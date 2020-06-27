package main

  type ListNode struct {
      Val int
      Next *ListNode
  }
func removeDuplicateNodes(head *ListNode) *ListNode {
    m := make(map[int]bool)
    var pre *ListNode
    root := head
    for head != nil {
        if m[head.Val] {
           pre.Next, head = head.Next , head.Next
        } else {
            m[head.Val] = true
            head, pre = head.Next, head
        }

    }
    return root
}