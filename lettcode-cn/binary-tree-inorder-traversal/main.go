package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * inorderTraversal 二叉树的中序遍历（迭代实现）
 * @param root 二叉树根节点
 * @return []int 中序遍历结果
 * @complexity 时间复杂度：O(n)，其中n是节点数
 * @complexity 空间复杂度：O(h)，其中h是树的高度
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	current := root

	for current != nil || len(stack) > 0 {
		// 将所有左子节点入栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// 弹出栈顶元素并访问
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// 处理右子树
		current = current.Right
	}

	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(root))
}
