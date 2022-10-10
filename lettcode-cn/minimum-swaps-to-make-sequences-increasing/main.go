package main

func minSwap(nums1, nums2 []int) int {
	n := len(nums1)
	a, b := 0, 1 //a不交换，b交换
	for i := 1; i < n; i++ {
		at, bt := a, b
		a, b = n, n
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			a = min(a, at)
			b = min(b, bt+1)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			a = min(a, bt)
			b = min(b, at+1)
		}
	}
	return min(a, b)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	println(minSwap([]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}))
}
