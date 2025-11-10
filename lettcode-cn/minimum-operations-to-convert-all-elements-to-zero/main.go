package main

import (
	"fmt"
)

// 官方解法：单调栈 O(n)
// 核心思想：
// 1. 从左到右遍历，维护一个单调递增栈
// 2. 栈中保存"活跃"的不同值（还未被完全处理的值）
// 3. 当遇到一个值时：
//    - 如果它比栈顶小，弹出栈顶（因为栈顶会被先处理掉）
//    - 如果栈为空或它比栈顶大，说明需要一次新操作
func minOperations(nums []int) int {
	operations := 0
	stack := []int{} // 单调递增栈

	for _, num := range nums {
		// 遇到0相当于一个分隔符，清空栈（重新开始新区域）
		if num == 0 {
			stack = stack[:0]
			continue
		}

		// 弹出所有 > num 的元素（注意：不弹出等于num的）
		// 为什么？因为这些更大的值会在处理num时一起被处理掉
		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}

		// 如果栈为空或栈顶 < num，说明需要新操作
		// 如果栈顶 == num，说明这个值已经在处理计划中，不需要额外操作
		if len(stack) == 0 || stack[len(stack)-1] < num {
			operations++
			stack = append(stack, num)
		}
	}

	return operations
}

// 测试用例结构
type testCase struct {
	name     string
	nums     []int
	expected int
}

func main() {
	// 表格驱动的测试用例
	testCases := []testCase{
		// 基础示例
		{
			name:     "示例1: [0,2]",
			nums:     []int{0, 2},
			expected: 1,
		},
		{
			name:     "示例2: [3,1,2,1]",
			nums:     []int{3, 1, 2, 1},
			expected: 3,
		},
		{
			name:     "示例3: [1,2,1,2,1,2]",
			nums:     []int{1, 2, 1, 2, 1, 2},
			expected: 4,
		},
		// 用户提到的关键测试
		{
			name:     "用户测试: [0,3,2,0]",
			nums:     []int{0, 3, 2, 0},
			expected: 2,
		},
		// 边界情况
		{
			name:     "全0数组",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "单元素非0",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "单元素0",
			nums:     []int{0},
			expected: 0,
		},
		// 更复杂的测试
		{
			name:     "复杂1: [1,2,3,0,3,2,1]",
			nums:     []int{1, 2, 3, 0, 3, 2, 1},
			expected: 6, // 1,2,3 各需要2次
		},
		{
			name:     "复杂2: [2,2,2,0,2,2]",
			nums:     []int{2, 2, 2, 0, 2, 2},
			expected: 2, // 被0分隔成两段
		},
		{
			name:     "复杂3: [1,2,1,0,1,2,1]",
			nums:     []int{1, 2, 1, 0, 1, 2, 1},
			expected: 4, // 1需要3次，2需要1次
		},
		// 连续相同值
		{
			name:     "连续相同: [5,5,5,5]",
			nums:     []int{5, 5, 5, 5},
			expected: 1,
		},
		// 交替模式
		{
			name:     "交替: [1,2,1,2,1]",
			nums:     []int{1, 2, 1, 2, 1},
			expected: 3, // 1需要2次，2需要1次
		},
		// 大值测试
		{
			name:     "大值: [100,50,100,50,100]",
			nums:     []int{100, 50, 100, 50, 100},
			expected: 4, // 50需要1次将其变0，然后100被分成3段，每段1次
		},
		// 用户提到的基础情况：无重复数字
		{
			name:     "无重复数字: [1,2,3,4,5]",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5, // 每个数字出现一次，需要5次操作
		},
	}

	fmt.Println("运行测试...")

	allPassed := true
	failedCases := []string{}

	for _, tc := range testCases {
		result := minOperations(tc.nums)
		passed := result == tc.expected
		if !passed {
			allPassed = false
			failedCases = append(failedCases,
				fmt.Sprintf("❌ %s: 输入=%v, 期望=%d, 实际=%d",
					tc.name, tc.nums, tc.expected, result))
		}
	}

	if allPassed {
		fmt.Printf("✅ 所有 %d 个测试用例通过\n", len(testCases))
	} else {
		fmt.Printf("\n❌ %d/%d 个测试失败：\n", len(failedCases), len(testCases))
		for _, msg := range failedCases {
			fmt.Println(msg)
		}
	}

	// 性能测试
	fmt.Println("\n=== 性能测试 ===")
	
	// 大数组测试
	bigArray := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		bigArray[i] = i % 100 + 1 // 100个不同的值，避免0
	}
	
	fmt.Printf("大数组测试 (长度: %d): ", len(bigArray))
	result := minOperations(bigArray)
	fmt.Printf("结果: %d\n", result)
	
	// 更大数组测试
	hugeArray := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		hugeArray[i] = i % 1000 + 1 // 1000个不同的值
	}
	
	fmt.Printf("超大数组测试 (长度: %d): ", len(hugeArray))
	result = minOperations(hugeArray)
	fmt.Printf("结果: %d\n", result)
	
	fmt.Println("\n✅ 性能测试完成 - 算法应该能在合理时间内处理大规模输入")
}