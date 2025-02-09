package main

import "fmt"

/*
* largestRectangleArea 计算柱状图中最大的矩形面积
* 算法思想:
* 1. 使用单调栈来解决,栈中存储柱子的下标
* 2. 栈中元素对应的高度保持单调递增,这样可以确保:
*    - 当遇到一个较小的高度时,可以确定栈顶元素向右扩展的边界
*    - 栈中下一个元素就是左边界
* 3. 遍历每个柱子:
*    - 如果栈空或当前高度大于栈顶高度,入栈
*    - 否则,不断弹出栈顶并计算面积,直到满足单调性
* 4. 最后处理栈中剩余元素,这些元素可以向右扩展到数组末尾
* 
* 时间复杂度: O(n) - 每个元素最多入栈出栈各一次
* 空间复杂度: O(n) - 栈的大小
*/
func largestRectangleArea(heights []int) int {
	// 初始化一个栈
	stack := []int{}
	// 初始化最大面积
	maxArea := 0

	// 遍历高度数组
	for i := 0; i < len(heights); i++ {
		// 如果栈为空或者当前高度大于栈顶高度,将当前高度入栈
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			// 如果当前高度小于栈顶高度,计算面积
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				// 弹出栈顶元素
				idx := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// 计算面积:宽度为右边界(i)减左边界(栈顶),高度为当前柱子高度
				width := i
				if len(stack) > 0 {
					width = i - stack[len(stack)-1] - 1
				}
				area := width * heights[idx]
				maxArea = max(maxArea, area)
			}
			stack = append(stack, i)
		}
	}

	// 计算剩余栈中的面积(这些柱子可以扩展到数组末尾)
	for len(stack) > 0 {
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}
		area := width * heights[idx]
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 输出: 10
}
