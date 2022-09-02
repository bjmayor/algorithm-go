package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 在左半边
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else {
			// 在右边
			if nums[len(nums)-1] >= target && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	return -1
}

func main() {
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 4)) // 4
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) //-1
	println(search([]int{1}, 0))                   //-1
	println(search([]int{5, 1, 3}, 2))             //-1

}
