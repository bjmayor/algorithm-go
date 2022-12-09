package main

func check(nums []int) bool {
	n := len(nums)
	nums = append(nums, nums...)
	for i := 0; i < n; i++ {
		ok := true
		for j := i; j < i+n-1; j++ {
			if nums[j] > nums[j+1] {
				ok = false
				break
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func main() {
	println(check([]int{3, 4, 5, 1, 2})) //true
	println(check([]int{2, 1, 3, 4}))    //false
	println(check([]int{1, 2, 3}))       //true
}
