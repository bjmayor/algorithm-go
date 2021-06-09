package main

import "fmt"

type ListNode struct {
      Val int
      Next *ListNode
  }
func removeElements(head *ListNode, val int) *ListNode {
    var nHead *ListNode = &ListNode{
        Val:  0,
        Next: nil,
    } 
    var pre *ListNode = nHead
    for head !=  nil {
    	if head.Val != val {
                pre.Next = head
                pre = head

        } else {
            pre.Next = nil
        }
        head = head.Next
    }
    return nHead.Next
}

func MakeListNode(data []int) *ListNode  {
    var head *ListNode = nil
    var pre *ListNode = nil
   for _, v := range data {
      if head == nil {
          head = &ListNode{
              Val:  v,
          }
          pre = head
      }  else {
          pre.Next = &ListNode{
              Val: v,
          }
          pre = pre.Next
      }
   }
   return head
}

func main()  {
    //h := MakeListNode([]int{1,2,6,3,4,5,6})
    //n :=removeElements(h, 6)
    //h := MakeListNode([]int{7,7,7,7})
    //n :=removeElements(h, 7)

    h := MakeListNode([]int{7,6,3,7})
    n :=removeElements(h, 7)
    fmt.Println(n)
}