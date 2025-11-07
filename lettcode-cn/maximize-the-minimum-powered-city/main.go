package main

import "fmt"

func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	
	// 用滑动窗口高效计算初始电量 O(n)
	power := make([]int64, n)
	// 初始化城市0的电量
	for i := 0; i <= r && i < n; i++ {
		power[0] += int64(stations[i])
	}
	left := power[0]
	right := power[0]
	
	// 滑动窗口计算其他城市的电量
	for j := 1; j < n; j++ {
		power[j] = power[j-1]
		// 加入右边界
		if j+r < n {
			power[j] += int64(stations[j+r])
		}
		// 移除左边界
		if j-r-1 >= 0 {
			power[j] -= int64(stations[j-r-1])
		}
		if power[j] > right {
			right = power[j]
		}
		if power[j] < left {
			left = power[j]
		}
	}
	
	// right 设为最大可能值
	right += int64(k)
	
	// 定义 check 函数（嵌套函数）
	var check func(int64) bool
	check = func(minPower int64) bool {
		add := make([]int64, n)
		power := int64(0)
		used := int64(0)
		
		// 初始化城市 0 的电量
		for j := 0; j < min(n, r+1); j++ {
			power += int64(stations[j])
		}
		
		for i := 0; i < n; i++ {
			// 维护滑动窗口
			if i > 0 {
				if i-r-1 >= 0 {
					power -= int64(stations[i-r-1]) + add[i-r-1]
				}
				if i+r < n {
					power += int64(stations[i+r]) + add[i+r]
				}
			}
			
			// 电量不足，需要新建
			if power < minPower {
				need := minPower - power
				buildPos := min(n-1, i+r)
				add[buildPos] += need
				power += need
				used += need
				
				if used > int64(k) {
					return false
				}
			}
		}
		
		return true
	}
	
	// 二分搜索
	for left < right {
		mid := (left + right + 1) >> 1
		if check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	
	return left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxPower([]int{1, 2, 4, 5, 0}, 1, 2)) // expect 5
}
