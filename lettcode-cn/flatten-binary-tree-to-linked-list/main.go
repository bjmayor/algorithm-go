package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * flatten 将二叉树展开为单链表
 * @param root 二叉树根节点
 * @complexity 时间复杂度：O(n)，其中n是节点数
 * @complexity 空间复杂度：O(1)，使用迭代方式，不使用额外空间
 */
func flatten(root *TreeNode) {
	// 处理空树的情况
	if root == nil {
		return
	}

	current := root
	for current != nil {
		// 如果当前节点有左子树
		if current.Left != nil {
			// 找到左子树的最右节点
			predecessor := current.Left
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}

			// 将当前节点的右子树接到左子树的最右节点
			predecessor.Right = current.Right
			// 将左子树移到右边
			current.Right = current.Left
			// 将左子树置空
			current.Left = nil
		}
		// 继续处理下一个节点
		current = current.Right
	}
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(root)

	// 打印展开后的链表
	for root != nil {
		fmt.Printf("%d -> ", root.Val)
		root = root.Right
	}
	fmt.Println("nil")
}
