package main

import "fmt"

/*
* findMin 在旋转排序数组中找到最小值
* @param nums 旋转排序数组
* @return int 最小值
*
* 解题思路:
* 1. 使用二分查找法
* 2. 如果nums[mid] > nums[right],说明最小值在mid的右边
* 3. 否则,最小值在mid的左边或就是mid
 */
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
