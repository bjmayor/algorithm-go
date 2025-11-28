package main

import (
	"fmt"
	"math"
	"sort"
)

func maxSumDivThree(nums []int) int {
	sort.Ints(nums) // 升序排序
	total := 0
	for _, v := range nums {
		total += v
	}
	r := total % 3
	if r == 0 {
		return total
	}
	// 收集模1和模2的元素
	var mod1 []int
	var mod2 []int
	for _, v := range nums {
		if v%3 == 1 {
			mod1 = append(mod1, v)
		} else if v%3 == 2 {
			mod2 = append(mod2, v)
		}
	}
	if r == 1 {
		// 移除一个模1或两个模2
		cand1 := math.MaxInt32
		if len(mod1) > 0 {
			cand1 = mod1[0]
		}
		cand2 := math.MaxInt32
		if len(mod2) >= 2 {
			cand2 = mod2[0] + mod2[1]
		}
		minRemove := min(cand1, cand2)
		if minRemove == math.MaxInt32 {
			return 0
		}
		return total - minRemove
	} else { // r == 2
		// 移除一个模2或两个模1
		cand1 := math.MaxInt32
		if len(mod2) > 0 {
			cand1 = mod2[0]
		}
		cand2 := math.MaxInt32
		if len(mod1) >= 2 {
			cand2 = mod1[0] + mod1[1]
		}
		minRemove := min(cand1, cand2)
		if minRemove == math.MaxInt32 {
			return 0
		}
		return total - minRemove
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 示例1
	nums1 := []int{3, 6, 5, 1, 8}
	fmt.Println(maxSumDivThree(nums1)) // 18

	// 示例2
	nums2 := []int{4}
	fmt.Println(maxSumDivThree(nums2)) // 0

	// 示例3
	nums3 := []int{1, 2, 3, 4, 4}
	fmt.Println(maxSumDivThree(nums3)) // 12
}
