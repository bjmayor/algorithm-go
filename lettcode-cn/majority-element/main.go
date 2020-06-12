package main

// hashmap 解法
func majorityElement1(nums []int) int {
	m := make(map[int]int, 0)
	for _,v := range nums {
		m[v]++
	}
	max := 0
	result := nums[0]
	for k,v := range m {
		if v > max {
			max = v
			result = k
		}
	}
	return result
}
// 利用众数 > 2/n 的特性。如果有解，连续长度一定最大
func majorityElement(nums []int) int {
	var res int
	count := 0
	for _, num := range nums {
		if count == 0 {
			res = num
		}
		if res == num {
			count++
		}else{
			count--
		}
	}
	return res
}