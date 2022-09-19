package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var ans int
	for head != nil {
		ans = ans*2 + head.Val
		head = head.Next
	}
	return ans
}
