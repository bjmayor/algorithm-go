package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		m[v] = i
	}
	stack := []int{}
	for j := len(nums2) - 1; j >= 0; j-- {
		for len(stack) > 0 && nums2[j] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		n := -1
		if len(stack) > 0 {
			n = stack[len(stack)-1]
		}
		if idx, ok := m[nums2[j]]; ok {
			ans[idx] = n
		}
		stack = append(stack, nums2[j])
	}
	return ans
}

func main() {

}
