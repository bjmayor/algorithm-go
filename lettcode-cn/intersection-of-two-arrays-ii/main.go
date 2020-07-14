package  main

import "fmt"

//小的放在前面
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}
	var res []int
	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}

func main()  {
	fmt.Println(intersect([]int{1,2,2,1}, []int{2,2})) //2,2
	fmt.Println(intersect([]int{4,9,5}, []int{9,4,9,8,4})) //2,2
}


