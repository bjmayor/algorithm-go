package main

func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minR := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > minR {
			return false
		}
		if nums[i-1] < minR {
			minR = nums[i-1]
		}
	}
	return true
}

func main() {
	println(isIdealPermutation([]int{1, 0, 2})) // true
	println(isIdealPermutation([]int{1, 2, 0})) // false
	println(isIdealPermutation([]int{0, 2, 1})) // true
	println(isIdealPermutation([]int{2, 0, 1})) // false
}
