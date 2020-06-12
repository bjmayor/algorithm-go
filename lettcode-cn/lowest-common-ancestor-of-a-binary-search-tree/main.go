package main

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
//递归解法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}


//非递归解法
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	for {
		if (p.Val < root.Val && root.Val < q.Val) || (q.Val < root.Val && p.Val > root.Val) {
			return root
		}
		if p.Val< root.Val && q.Val < root.Val {
			root = root.Left
		}

		if p.Val> root.Val && q.Val > root.Val {
			root = root.Right
		}
	}
}