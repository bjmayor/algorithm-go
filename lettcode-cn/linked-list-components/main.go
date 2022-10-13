package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	m := map[int]struct{}{}
	for _, v := range nums {
		m[v] = struct{}{}
	}
	ans := 0
	inSet := false
	for head != nil {
		if _, ok := m[head.Val]; ok {
			if !inSet {
				inSet = true
				ans++
			}
		} else {
			inSet = false
		}
		head = head.Next
	}
	return ans
}
