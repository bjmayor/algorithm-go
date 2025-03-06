package main

import "fmt"

/*
*
关键是异或和为 0 的子数组为美丽子数组
*/
func beautifulSubarrays(nums []int) int64 {
	m := make(map[int]int64)
	m[0] = 1
	var ans int64 = 0
	preXor := 0
	for _, v := range nums {
		preXor ^= v
		ans += m[preXor]
		m[preXor]++
	}
	return ans
}

func main() {
	nums := []int{4, 3, 1, 2, 4}
	fmt.Println(beautifulSubarrays(nums))
}
