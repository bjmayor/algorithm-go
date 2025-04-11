package main

/*
1123. 最深叶节点的最近公共祖先
给你一个有根节点 root 的二叉树，返回它 最深的叶节点的最近公共祖先 。

回想一下：

叶节点 是二叉树中没有子节点的节点
树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
如果我们假定 A 是一组节点 S 的 最近公共祖先，S 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lcaDeepestLeaves 返回最深叶节点的最近公共祖先
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var result *TreeNode
	maxDepth := 0

	// 后序遍历函数
	var postorder func(node *TreeNode, depth int) int
	postorder = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth
		}

		// 后序遍历：左子树 -> 右子树 -> 根节点
		leftDepth := postorder(node.Left, depth+1)
		rightDepth := postorder(node.Right, depth+1)

		// 如果左右子树深度相同且等于最大深度，更新结果
		if leftDepth == rightDepth && leftDepth >= maxDepth {
			maxDepth = leftDepth
			result = node
		}

		// 返回当前子树的最大深度
		if leftDepth > rightDepth {
			return leftDepth
		}
		return rightDepth
	}

	postorder(root, 0)
	return result
}

// createTree 根据数组创建二叉树，-1表示空节点
func createTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nums) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) && nums[i] != -1 {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != -1 {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func main() {
	// 测试用例1
	nums1 := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root1 := createTree(nums1)
	result1 := lcaDeepestLeaves(root1)
	println("测试用例1结果:", result1.Val) // 应该输出 2

	// 测试用例2
	nums2 := []int{1}
	root2 := createTree(nums2)
	result2 := lcaDeepestLeaves(root2)
	println("测试用例2结果:", result2.Val) // 应该输出 1

	// 测试用例3
	nums3 := []int{0, 1, 3, -1, 2}
	root3 := createTree(nums3)
	result3 := lcaDeepestLeaves(root3)
	println("测试用例3结果:", result3.Val) // 应该输出 2
}
