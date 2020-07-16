package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type IndexNode struct {
	Idx int
	Val int
}
func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return nil
	}
	n:=1
	c :=  head
	for c.Next != nil {
		c = c.Next
		n++
	}
	var res   = make([]int, n)
	i := 0
	c = head
	s := make([]*IndexNode, 0)
	s = append(s, &IndexNode{i, c.Val})
	for c.Next != nil {
		c = c.Next
		i++
		for len(s) > 0 {
			top := s[len(s)-1]
			if top.Val < c.Val {
				res[top.Idx] = c.Val
				s = s[:len(s)-1]
			} else {
				break
			}
		}
		s = append(s,  &IndexNode{i, c.Val})
	}
	return res
}

func main()  {
	n1 := &ListNode{Val:2}
	n2 := &ListNode{Val:1}
	n5 := &ListNode{Val:5}
	n1.Next = n2
	n2.Next = n5
	fmt.Println(nextLargerNodes(n1))
}