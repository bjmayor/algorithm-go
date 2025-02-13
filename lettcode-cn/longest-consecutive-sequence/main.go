package main

import "fmt"

/*
https://leetcode.cn/problems/longest-consecutive-sequence/

解题思路:
1. 使用哈希表存储所有数字,实现O(1)时间复杂度的查找
2. 遍历哈希表中的每个数字k
3. 如果k-1存在于哈希表中,说明k不是连续序列的起点,跳过
4. 如果k-1不存在,说明k是一个连续序列的起点:
   - 从k开始向后查找k+1,k+2等数是否存在
   - 记录当前连续序列的长度
   - 更新最长连续序列的长度
5. 最终返回最长连续序列的长度

时间复杂度: O(n) - 每个数字最多被访问两次
空间复杂度: O(n) - 需要哈希表存储所有数字
*/

func longestConsecutive(nums []int) int {
	// 使用map存储所有数字,值使用空结构体节省内存
	m := make(map[int]struct{}, 0)
	var result = 0
	// 处理特殊情况:数组不为空时,最小长度为1
	if len(nums) > 0 {
		result = 1
	}
	// 将所有数字存入map
	for _, k := range nums {
		m[k] = struct{}{}
	}
	// 遍历map中的每个数字
	for k := range m {
		// 如果k-1存在,说明k不是连续序列的起点
		if _, ok := m[k-1]; ok {
			continue
		}
		// k是连续序列的起点,计算以k开始的连续序列长度
		len := 1
		detect := k + 1
		// 不断检查下一个数字是否存在
		for {
			if _, ok := m[detect]; !ok {
				break
			}
			len++
			detect++
		}
		// 更新最长连续序列的长度
		if len > result {
			result = len
		}
	}

	return result
}

func main() {
	// 测试用例
	tests := []struct {
		Nums     []int
		Expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4}, // 连续序列: 1,2,3,4
		{[]int{0}, 1},                     // 单个数字
		{[]int{1, 3, 5, 2, 4}, 5},        // 连续序列: 1,2,3,4,5
	}
	// 运行测试
	for _, t := range tests {
		var real = longestConsecutive(t.Nums)
		if real != t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real)
		}
	}
}
