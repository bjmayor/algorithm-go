package main

import "fmt"

/*
https://leetcode.cn/problems/sort-colors/?envType=study-plan-v2&envId=top-100-liked
75. 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

/*
* sortColors 对颜色进行分类
* 算法思想:
* 1. 使用三个指针: low, mid, high
* 2. 遍历数组, 根据当前元素的值进行交换和指针移动
* 3. 最终实现原地排序
*
* 时间复杂度: O(n) - 单次遍历
* 空间复杂度: O(1) - 只使用常数空间
 */
func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
