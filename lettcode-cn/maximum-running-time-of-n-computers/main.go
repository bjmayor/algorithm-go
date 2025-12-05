package main

import "fmt"

// maxRunTime 返回 n 台电脑同时运行的最长分钟数
// 思路：二分答案 + 判定函数
func maxRunTime(n int, batteries []int) int64 {
	var sum int64
	for _, b := range batteries {
		sum += int64(b)
	}
	low, high := int64(0), sum/int64(n)

	// check 判断能否让 n 台电脑同时运行 t 分钟
	check := func(t int64) bool {
		var total int64
		for _, b := range batteries {
			if int64(b) >= t {
				total += t
			} else {
				total += int64(b)
			}
			// 提前返回优化
			if total >= t*int64(n) {
				return true
			}
		}
		return total >= t*int64(n)
	}

	// 二分找最大可行值
	for low < high {
		mid := (low + high + 1) / 2
		if check(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func main() {
	// 示例测试
	fmt.Println(maxRunTime(2, []int{3, 3, 3}))    // 4
	fmt.Println(maxRunTime(2, []int{1, 1, 1, 1})) // 2
	fmt.Println(maxRunTime(1, []int{10}))         // 10
}
