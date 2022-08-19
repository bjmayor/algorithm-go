package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums11 := make([]int, m)
	copy(nums11, nums1)
	pos, i, j := 0, 0, 0
	for i, j = 0, 0; i < m && j < n; {
		if nums11[i] <= nums2[j] {
			nums1[pos] = nums11[i]
			i++
			pos++
		} else {
			nums1[pos] = nums2[j]
			j++
			pos++
		}
	}
	if i < m {
		copy(nums1[i+n:], nums11[i:m])
	}
	if j < n {
		copy(nums1[j+m:], nums2[j:n])
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
