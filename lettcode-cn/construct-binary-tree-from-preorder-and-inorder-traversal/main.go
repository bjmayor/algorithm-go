package main

import "fmt"

/**
 * TreeNode 二叉树节点的数据结构定义
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * buildTree 根据前序遍历和中序遍历构建二叉树
 * @param preorder 前序遍历序列
 * @param inorder 中序遍历序列
 * @return *TreeNode 构建的二叉树根节点
 * @complexity 时间复杂度：O(n)，其中n是节点数
 * @complexity 空间复杂度：O(n)，递归调用栈的深度
 *
 * 算法思想:
 * 1. 前序遍历的第一个节点是根节点
 * 2. 在中序遍历中找到根节点的位置，可以将数组分成左右子树
 * 3. 递归构建左右子树
 *    - 左子树：前序遍历[1:rootIndex+1], 中序遍历[:rootIndex]
 *    - 右子树：前序遍历[rootIndex+1:], 中序遍历[rootIndex+1:]
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 处理空树的情况
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 前序遍历的第一个节点是根节点
	root := &TreeNode{Val: preorder[0]}

	// 在中序遍历中找到根节点的位置
	rootIndex := 0
	for i, v := range inorder {
		if v == preorder[0] {
			rootIndex = i
			break
		}
	}

	// 递归构建左右子树
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

/**
 * main 主函数，用于测试二叉树的构建和遍历
 */
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)

	// 打印树的结构
	fmt.Println("构建的二叉树结构如下:")
	fmt.Println("     ", root.Val)
	fmt.Println("    /  \\")
	fmt.Println("   " + fmt.Sprint(root.Left.Val) + "    " + fmt.Sprint(root.Right.Val))
	fmt.Println("        /  \\")
	fmt.Println("       " + fmt.Sprint(root.Right.Left.Val) + "   " + fmt.Sprint(root.Right.Right.Val))

	// 前序遍历
	fmt.Print("前序遍历: ")
	preorderTraversal(root)
	fmt.Println()

	// 中序遍历
	fmt.Print("中序遍历: ")
	inorderTraversal(root)
	fmt.Println()
}

/**
 * preorderTraversal 前序遍历二叉树
 * 遍历顺序：根节点 -> 左子树 -> 右子树
 */
func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

/**
 * inorderTraversal 中序遍历二叉树
 * 遍历顺序：左子树 -> 根节点 -> 右子树
 */
func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Print(root.Val, " ")
	inorderTraversal(root.Right)
}
