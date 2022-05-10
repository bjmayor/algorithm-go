package main

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
func isSymmetric(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue,root)
	for len(queue) > 0 {
		size:=len(queue)
		for i:=0;i<=size/2;i++ {
			head := queue[i]
			tail := queue[size-1-i]
			if head!=nil && tail==nil {
				return false
			}
			if head==nil && tail!=nil {
				return false
			}
			if head!=nil && tail != nil && head.Val != tail.Val {
				return false
			}
		}
		for i:=0;i<=size-1;i++ {
			n := queue[i]
			if n != nil {
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			}
		}
		queue = queue[size:]
	}
	return true
}

func main()  {
	nodes:= []*TreeNode{}
	//vs := []interface{}{1,2,2,nil,3,nil,3};
	vs := []interface{}{1,2,2,3,4,4,3};
	for i:=0;i<len(vs);i++ {
		if vs[i] == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &TreeNode{Val: vs[i].(int)})
		}
	}
	for i:=0;i<len(vs)/2;i++ {
		left :=  i*2+1
		right := i*2+2
		nodes[i].Left = nodes[left]
		nodes[i].Right= nodes[right]
	}
	println(isSymmetric(nodes[0]))
}