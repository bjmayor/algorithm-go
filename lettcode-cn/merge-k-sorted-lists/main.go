package main

type ListNode struct {
	Val int
	Next *ListNode
}
/*
//优先队列合并 40ms
func mergeKLists(lists []*ListNode) *ListNode {

	//把nil放到最后面
	k := len(lists)-1//k 之后都是无效值
	i:=0
	for i<=k {
		if lists[k] == nil {
			k--
			if k < i {
				break
			}
		}
		if lists[i] ==  nil {
			lists[i], lists[k] = lists[k], lists[i]
			k--
		} else {
			i++
		}
	}
	lists = lists[:k+1]
	if len(lists) == 0 {
		return nil
	}
	var head *ListNode
	var cur *ListNode
	first := true
	for len(lists) > 1 {
		if lists[0] == nil {
			lists = lists[1:]
			continue
		}
		if first {
			for i:=1;i<len(lists);i++ {
				tmp := lists[i]
				for j:=i-1;j>=0;j-- {
					if lists[j].Val>tmp.Val {
						lists[j+1] = lists[j]
						if j == 0 {
							lists[0] = tmp
						}
					} else {
						lists[j+1] = tmp
						break
					}
				}
			}
			first = false
		} else {
			idx :=0
			for i:=1;i<len(lists);i++ {
				if lists[i].Val < lists[idx].Val {
					lists[i], lists[idx] = lists[idx], lists[i]
					idx = i
				} else {
					break
				}
			}
		}

		s := lists[0]
		//去掉他
		if s.Next == nil {
			lists = lists[1:]
		} else {
			lists[0] = s.Next
		}
		if head == nil {
			head = s
			cur = s
		} else {
			cur.Next = s
			cur = s
		}
	}
	if cur != nil {
		cur.Next = lists[0]
	} else {
		head = lists[0]
	}
	return head
}
*/

//分治法12ms
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)== 0 {
		return nil
	}
	if len(lists)== 1 {
		return lists[0]
	}
	mid := len(lists)/2
	l1 := mergeKLists(lists[:mid])
	l2 := mergeKLists(lists[mid:])
	return mergeTwoLists(l1, l2)
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := l1
	h2 := l2

	var head *ListNode
	var cur *ListNode
	for h1!=nil && h2!=nil {
		s := h1
		if h2.Val < h1.Val {
			s = h2
			h2 = h2.Next
		} else {
			h1 = h1.Next
		}
		if head == nil {
			head = s
			cur = s
		}
		cur.Next, cur = s, s
	}
	if h1!=nil {
		if cur != nil {
			cur.Next = h1
		} else {
			head = h1
		}
	}
	if h2!=nil {
		if cur != nil {
			cur.Next = h2
		} else {
			head = h2
		}
	}
	return head
}

func main()  {
	//n11 := ListNode{Val:2}
	//n14 := ListNode{Val:4}
	//n15 := ListNode{Val:5}
	//n11.Next = &n14
	//n14.Next = &n15

	//n21 := ListNode{Val:-1}
	//n23 := ListNode{Val:3}
	//n24 := ListNode{Val:4}
	//n21.Next = &n23
	//n23.Next = &n24

	//n32 := ListNode{Val:2}
	//n36 := ListNode{Val:6}
	//n32.Next = &n36

	mergeKLists([]*ListNode{
		 nil,
	})
}
