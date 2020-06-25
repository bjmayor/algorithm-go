package main
func removeElement(nums []int, val int) int {
	last := len(nums)-1 // 最后一个非 val, 替换完就变成了val
	for i:=len(nums)-1;i>=0;i-- {
		if nums[i] == val {
			nums[i], nums[last] = nums[last], nums[i]
			last--
		}
	}
	return last+1
}
