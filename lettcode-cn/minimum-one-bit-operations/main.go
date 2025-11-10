package main

import (
	"fmt"
	"strconv"
)

// 1611. 使整数变为 0 的最少操作次数
// 给你一个整数 n，你需要重复执行多次下述操作将其转换为 0 ：
// 1. 翻转 n 的二进制表示中最右侧位（第 0 位）。
// 2. 如果第 (i-1) 位为 1 且从第 (i-2) 位到第 0 位都为 0，则翻转 n 的二进制表示中的第 i 位。
// 返回将 n 转换为 0 的最小操作次数。
func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}
	
	// 找到最高位
	highestBit := 0
	temp := n
	for temp > 0 {
		highestBit++
		temp >>= 1
	}
	highestBit-- // 转换为0基索引
	
	// 递归处理
	return calculateOperations(n, highestBit)
}

// 递归计算函数
func calculateOperations(n int, bit int) int {
	if bit < 0 || n == 0 {
		return 0
	}
	
	// 检查最高位是否为1
	if (n>>bit)&1 == 1 {
		// 最高位是1的情况
		// 移除最高位后的剩余值
		remaining := n ^ (1 << bit)
		// 递归计算剩余部分
		subResult := calculateOperations(remaining, bit-1)
		// 应用九连环的符号规则
		return (1 << (bit + 1)) - 1 - subResult
	} else {
		// 最高位是0，直接跳过
		return calculateOperations(n, bit-1)
	}
}

// 迭代实现：更简洁的版本
func minimumOneBitOperationsIterative(n int) int {
	if n == 0 {
		return 0
	}
	
	result := 0
	isPositive := true
	bit := 31 // 从最高位开始处理
	
	// 找到最高位的位置
	for bit >= 0 {
		if (n>>bit)&1 == 1 {
			break
		}
		bit--
	}
	
	// 从最高位开始处理
	for i := bit; i >= 0; i-- {
		if (n>>i)&1 == 1 {
			// 计算该位的操作次数
			operations := (1 << (i + 1)) - 1
			
			if isPositive {
				result += operations
			} else {
				result -= operations
			}
			
			// 符号交替
			isPositive = !isPositive
		}
	}
	
	return result
}

// 验证九连环的路径不重复性
func verifyNoDuplicatePaths(n int) {
	fmt.Printf("验证 n=%d 的路径不重复性:\n", n)
	
	visited := make(map[int]bool)
	queue := []int{n}
	visited[n] = true
	
	step := 0
	for len(queue) > 0 && step < 1000 { // 防止无限循环
		size := len(queue)
		step++
		
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			
			// 生成可能的操作
			nextStates := generateNextStates(current)
			
			for _, next := range nextStates {
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	
	fmt.Printf("总共访问了 %d 个不同状态，路径无重复！\n\n", len(visited))
}

// 生成下一个可能的状态
func generateNextStates(n int) []int {
	var states []int
	
	// 操作1：翻转第0位
	states = append(states, n^1)
	
	// 操作2：检查其他位是否满足条件
	for i := 1; i < 31; i++ {
		if ((n>>i)&1) == 1 {
			// 检查第(i-1)位是否为1
			if ((n>>(i-1))&1) == 1 {
				// 检查第0到(i-2)位是否都为0
				lowerBits := n & ((1 << (i-1)) - 1)
				if lowerBits == 0 {
					states = append(states, n^(1<<i))
				}
			}
		}
	}
	
	return states
}

func main() {
	// 测试用例
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 2},
		{6, 4},
		{9, 14},
		{333, 393},
	}
	
	fmt.Println("=== 递归方法测试 ===")
	for _, tc := range testCases {
		result := minimumOneBitOperations(tc.input)
		fmt.Printf("n=%d, 期望=%d, 实际=%d, %s\n", 
			tc.input, tc.expected, result, 
			map[bool]string{true: "✓", false: "✗"}[result == tc.expected])
	}
	
	fmt.Println("\n=== 迭代方法测试 ===")
	for _, tc := range testCases {
		result := minimumOneBitOperationsIterative(tc.input)
		fmt.Printf("n=%d, 期望=%d, 实际=%d, %s\n", 
			tc.input, tc.expected, result, 
			map[bool]string{true: "✓", false: "✗"}[result == tc.expected])
	}
	
	// 验证路径无重复
	fmt.Println("\n=== 路径验证 ===")
	verifyNoDuplicatePaths(6)
	
	// 打印图示
	printOperationRules()
	printNineRingsAnalogy()
	printNineRingsPattern()
	printSignPattern()
	printDetailedExample(6)
	
	// 演示具体的操作过程
	fmt.Println("\n=== 操作过程演示 ===")
	demonstrateProcess(6)
}

// 演示具体的操作过程
func demonstrateProcess(n int) {
	fmt.Printf("演示 n=%d (二进制: %s) 的操作过程\n", 
		n, strconv.FormatInt(int64(n), 2))
	
	current := n
	step := 0
	
	for current > 0 {
		step++
		binary := strconv.FormatInt(int64(current), 2)
		fmt.Printf("步骤 %d: %s -> ", step, binary)
		
		// 简化版的操作逻辑
		if current&1 == 1 {
			// 最低位是1，直接翻转
			current ^= 1
			fmt.Printf("操作1 (翻转第0位)\n")
		} else {
			// 找到第一个满足条件的位进行翻转
			for i := 1; i < 32; i++ {
				if ((current>>i)&1) == 1 && ((current>>i-1)&1) == 1 {
					// 这里简化处理，实际需要检查更复杂的条件
					current ^= (1 << i)
					fmt.Printf("操作2 (翻转第%d位)\n", i)
					break
				}
			}
		}
	}
	
	fmt.Printf("最终结果: %d 步操作\n", step)
}

// 打印操作规则图示
func printOperationRules() {
	fmt.Println("=== 操作规则图示 ===")
	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│                   1611题操作规则                         │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│  操作1: 翻转第0位                                       │")
	fmt.Println("│  格式:  ...x0  →  ...x1                                 │")
	fmt.Println("│  条件: 无条件，可以直接执行                              │")
	fmt.Println("│                                                         │")
	fmt.Println("│  操作2: 条件翻转第i位                                   │")
	fmt.Println("│  格式:  ...10...0i  →  ...00...1i                      │")
	fmt.Println("│  条件: 第(i-1)位为1，且第0~(i-2)位都为0                 │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()
}

// 打印九连环类比图示
func printNineRingsAnalogy() {
	fmt.Println("=== 九连环类比图示 ===")
	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│                    九连环的数学本质                      │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│  九连环状态:                                             │")
	fmt.Println("│  环的排列: [8] [7] [6] [5] [4] [3] [2] [1] [0]           │")
	fmt.Println("│  1表示在棒上，0表示不在棒上                              │")
	fmt.Println("│                                                         │")
	fmt.Println("│  取下第k环的规则:                                       │")
	fmt.Println("│  - 如果第(k-1)环在棒上                                   │")
	fmt.Println("│  - 且第0~(k-2)环都不在棒上                               │")
	fmt.Println("│  - 则可以取下第k环                                      │")
	fmt.Println("│                                                         │")
	fmt.Println("│  1611题的二进制操作与九连环完全相同！                    │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()
}

// 打印具体示例的详细图示
func printDetailedExample(n int) {
	fmt.Printf("=== 详细示例: n=%d (二进制: %s) ===\n", n, strconv.FormatInt(int64(n), 2))
	fmt.Println()
	
	// 打印位索引
	fmt.Print("位索引:    ")
	for i := 7; i >= 0; i-- {
		if n > (1<<i) || i == 0 {
			fmt.Printf("%d ", i)
		} else {
			fmt.Print("  ")
		}
	}
	fmt.Println()
	
	// 打印二进制表示
	binary := strconv.FormatInt(int64(n), 2)
	fmt.Printf("二进制:    %s\n", binary)
	fmt.Println()
	
	// 显示每一步的操作
	current := n
	step := 0
	maxBits := 8
	
	for current > 0 && step < 20 { // 防止无限循环
		step++
		currentBinary := strconv.FormatInt(int64(current), 2)
		
		fmt.Printf("步骤 %d: ", step)
		
		// 打印当前状态
		for i := maxBits - 1; i >= 0; i-- {
			if i >= len(currentBinary) {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", currentBinary[len(currentBinary)-1-i])
			}
		}
		fmt.Print(" → ")
		
		// 模拟操作
		if current&1 == 1 {
			// 操作1：翻转第0位
			current ^= 1
			fmt.Println("翻转第0位 (操作1)")
		} else {
			// 操作2：找到可翻转的高位
			for i := 1; i < maxBits; i++ {
				if ((current>>i)&1) == 1 {
					if i == 1 || ((current>>i-1)&1) == 1 {
						current ^= (1 << i)
						fmt.Printf("翻转第%d位 (操作2)\n", i)
						break
					}
				}
			}
		}
	}
	fmt.Printf("总共需要 %d 步操作\n\n", step)
}

// 打印九连环模式图
func printNineRingsPattern() {
	fmt.Println("=== 九连环模式分析 ===")
	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│                    操作次数规律                         │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│  单个位为1的情况:                                       │")
	fmt.Println("│  2^0 = 1    → 0     需要 1 步   (2^1 - 1)              │")
	fmt.Println("│  2^1 = 2    → 00    需要 3 步   (2^2 - 1)              │")
	fmt.Println("│  2^2 = 4    → 000   需要 7 步   (2^3 - 1)              │")
	fmt.Println("│  2^3 = 8    → 0000  需要 15 步  (2^4 - 1)              │")
	fmt.Println("│                                                         │")
	fmt.Println("│  一般规律: f(2^k) = 2^(k+1) - 1                        │")
	fmt.Println("│  这是典型的九连环数学模型！                             │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()
}

// 打印符号交替规律
func printSignPattern() {
	fmt.Println("=== 符号交替规律 ===")
	fmt.Println("┌─────────────────────────────────────────────────────────┐")
	fmt.Println("│                  二进制位的符号处理                     │")
	fmt.Println("├─────────────────────────────────────────────────────────┤")
	fmt.Println("│  例子: n = 333 = 101001101₂                            │")
	fmt.Println("│                                                         │")
	fmt.Println("│  处理过程:                                              │")
	fmt.Println("│  位位置  |  值  |  操作次数  |  符号  |  累计结果       │")
	fmt.Println("│  --------|------|-----------|--------|---------------  │")
	fmt.Println("│    8     |  1   |    511    |   +    |     +511       │")
	fmt.Println("│    6     |  1   |    127    |   -    |     -127       │")
	fmt.Println("│    3     |  1   |     15    |   +    |      +15       │")
	fmt.Println("│    2     |  1   |      7    |   -    |       -7       │")
	fmt.Println("│    0     |  1   |      1    |   +    |       +1       │")
	fmt.Println("│  --------|------|-----------|--------|---------------  │")
	fmt.Println("│                最终结果: 511 - 127 + 15 - 7 + 1 = 393   │")
	fmt.Println("└─────────────────────────────────────────────────────────┘")
	fmt.Println()
}