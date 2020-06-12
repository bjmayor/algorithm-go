package main

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
// bfs 遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0, 0)
   queue := make([]*TreeNode,0,0)
   queue = append(queue, root)
   for len(queue) > 0 {
   		num := len(queue)
	   tmp := make([]int, 0, 0)
	   for i:=0;i<num;i++ {
		   node := queue[0]
		   tmp = append(tmp, node.Val)
		   queue = queue[1:]
		   if node.Left != nil {
			   queue = append(queue, node.Left)
		   }

		   if node.Right != nil {
			   queue = append(queue, node.Right)
		   }
	   }
	   result = append(result, tmp)
   }
   return result
}
 // dfs 遍历
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0, 0)
	dfs(root, 0, &result)
	return result
}

func dfs(root *TreeNode, level int, result *[][]int)  {
	if root == nil {
		return
	}
	if len(*result) < level+1 {
		*result = append(*result, make([]int, 0, 0))
	}
	tmp := (*result)[level]
	tmp = append(tmp, root.Val)
	(*result)[level] = tmp
	dfs(root.Left, level+1, result)
	dfs(root.Right, level+1, result)
}