package main

import "fmt"

/*
3396. 使数组元素互不相同所需的最少操作次数
*/
func minimumOperations(nums []int) int {
	m := make(map[int]bool)
	toDel := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = true
		} else {
			toDel = i + 1
			break
		}
	}
	if toDel%3 == 0 {
		return toDel / 3
	}
	return toDel/3 + 1
}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4, 2, 3, 3, 5, 7})) // 2
	fmt.Println(minimumOperations([]int{4, 5, 6, 4, 4}))             // 2
	fmt.Println(minimumOperations([]int{6, 7, 8, 9}))                // 0
}
