package main
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func bstFromPreorder(preorder []int) *TreeNode {
    var root *TreeNode
    for i, v := range preorder {
        if root == nil {
            root = &TreeNode{Val:v}
        } else {
            if v < root.Val {
                j:=i+1
                for j=i+1; j<len(preorder);j++ {
                    if preorder[j] > root.Val {
                       break
                    }
                }
                root.Left = bstFromPreorder(preorder[i:j])
                root.Right= bstFromPreorder(preorder[j:])
                return root
            } else {
                root.Right = bstFromPreorder(preorder[i:])
                return root
            }
        }
    }
    return root
}
