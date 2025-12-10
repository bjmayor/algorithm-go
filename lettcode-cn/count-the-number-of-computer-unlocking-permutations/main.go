package main

import "fmt"

const MOD = 1000000007

// countPermutations 返回能解锁所有计算机的排列数（模 MOD）。
// 若 complexity[0] 不是严格小于其余元素（即不是唯一的最小值），返回 0。
func countPermutations(complexity []int) int {
	n := len(complexity)
	if n == 0 {
		return 0
	}
	// 判断第0个是否为唯一最小值
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}
	// 计算 (n-1)! % MOD
	ans := 1
	for i := 2; i <= n-1; i++ {
		ans = (ans * i) % MOD
	}
	return ans
}

func main() {
	// 直接使用 topic.md 中的示例进行演示
	examples := [][]int{
		{1, 2, 3},
		{3, 3, 3, 4, 4, 4},
	}
	for _, ex := range examples {
		fmt.Printf("input: %v\n", ex)
		fmt.Printf("output: %d\n\n", countPermutations(ex))
	}
}
