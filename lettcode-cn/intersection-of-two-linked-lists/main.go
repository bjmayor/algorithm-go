package main


 type ListNode struct {
     Val int
     Next *ListNode
 }
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m := map[*ListNode]struct{}{}
	for headA!=nil {
	    m[headA] = struct{}{}
	    headA = headA.Next
    }
	for headB!=nil {
		if _, ok := m[headB];ok {
			return headB
		}
		headB = headB.Next
	}
    return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	for pa !=  pb {
		if pa == nil { //链接B
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil { //链接B
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}