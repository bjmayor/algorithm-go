package main

import "fmt"

/**
 * Definition for singly-linked list.

 */

 type ListNode struct {
	     Val int
	     Next *ListNode
 }


func nodeAdd(l *ListNode, num int) {
	if l != nil {
		sum := l.Val + num
		if sum>=10 {
			l.Val = sum-10
			if l.Next == nil {
				l.Next = new(ListNode)
			}
			nodeAdd(l.Next, 1)
		} else {
			l.Val = sum
		}
	}
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result  ListNode
	var current = &result
	for l1 != nil || l2 != nil {
		var sum int
		if l1!=nil {
			sum = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		nodeAdd(current, sum)
		if current.Next == nil {
			next := new(ListNode)
			current.Next = next
		}
		current = current.Next
	}
	return &result
}

func printNode(l *ListNode)  {
	for l!=nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Println()
}

func listLink(data []int) *ListNode {
	var l = new(ListNode)
	
	return l
}

func main()  {
	l1 := ListNode{
		Val:  9,
		Next: nil,
	}
	node8 := ListNode{
		Val:  9,
		Next: nil,
	}
	l1.Next = &node8

	l2 := ListNode{
		Val:  1,
		Next: nil,
	}

	printNode(addTwoNumbers(&l1,&l2))
}