package main

import "fmt"

/*

https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/submissions/597894754/?envType=daily-question&envId=2025-02-09

80. 删除有序数组中的重复项 II

给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

/*
* removeDuplicates 删除有序数组中的重复项,使每个元素最多出现两次
* 算法思想:
* 1. 使用快慢指针:
*    - slow 指向当前可写入位置
*    - fast 用于遍历数组
* 2. 判断条件:
*    - 当 nums[fast] != nums[slow-2] 时,说明当前元素可以写入
*    - 将 fast 位置的元素写入 slow 位置
*    - slow 向前移动
* 3. 最终 slow 即为新数组的长度
*
* 时间复杂度: O(n) - n为数组长度
* 空间复杂度: O(1) - 原地修改数组,仅使用常数额外空间
 */
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	
	slow := 2 // 从第三个位置开始判断
	for fast := 2; fast < n; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
