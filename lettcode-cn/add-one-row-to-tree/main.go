package main


  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    if depth==1 {
        n := new(TreeNode)
        n.Val = val
        n.Left = root
        return n
    }
    q := []*TreeNode{root}
    for level:=1;len(q)>0 && level<=depth-2;level++ {
        size := len(q)
        for i:=0;i<size;i++ {
            n := q[i]
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
        q = q[size:]
    }
    for _, n := range q {
        nl := new(TreeNode)
        nl.Val = val
        nl.Left = n.Left
        n.Left = nl
        nr := new(TreeNode)
        nr.Val = val
        nr.Right= n.Right
        n.Right= nr
    }
    return root
}

func main()  {
   n4 := &TreeNode{
       Val:   4,
       Left:  nil,
       Right: nil,
   }

    n3 := &TreeNode{
        Val:   3,
        Left:  nil,
        Right: nil,
    }

    n2 := &TreeNode{
        Val:   2,
        Left:  n4,
        Right: nil,
    }

    n1 := &TreeNode{
        Val:   1,
        Left:  n2,
        Right: n3,
    }

    addOneRow(n1, 5, 4)
}
