package main

import (
	"container/list"
	"fmt"
)

/**
 * @title 2612. 最少翻转操作数
 * @description 计算将数组中唯一的1移动到每个位置所需的最少翻转次数
 * @param n {int} 数组长度
 * @param p {int} 初始1的位置
 * @param banned {[]int} 禁止访问的位置数组
 * @param k {int} 翻转子数组的长度
 * @return {[]int} 返回移动到每个位置所需的最少翻转次数数组
 */
func minReverseOperations(n int, p int, banned []int, k int) []int {
	// 初始化答案数组
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}

	// 创建禁止访问位置的集合
	bannedSet := make(map[int]bool)
	for _, b := range banned {
		bannedSet[b] = true
	}

	// 使用两个平衡树分别存储奇偶下标
	even := NewBalancedSet()
	odd := NewBalancedSet()

	// 初始化可访问的位置
	for i := 0; i < n; i++ {
		if !bannedSet[i] && i != p {
			if i%2 == 0 {
				even.Insert(i)
			} else {
				odd.Insert(i)
			}
		}
	}

	// BFS队列
	queue := list.New()
	queue.PushBack(p)
	ans[p] = 0

	// 广度优先搜索
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)

		// 计算当前位置可以翻转到的范围
		minPos := max(cur-k+1, k-cur-1)     // max(i-k+1, k-i-1)
		maxPos := min(cur+k-1, 2*n-k-cur-1) // min(i+k-1, 2n-k-i-1)
		minPos = max(0, minPos)             // 确保不小于0
		maxPos = min(n-1, maxPos)           // 确保不超过n-1

		// 选择对应的集合
		set := even
		if minPos%2 == 1 {
			set = odd
		}

		// 遍历可能的新位置（按照奇偶性递增）
		for newPos := minPos; newPos <= maxPos; newPos += 2 {
			if set.Contains(newPos) {
				set.Remove(newPos)
				ans[newPos] = ans[cur] + 1
				queue.PushBack(newPos)
			}
		}
	}

	return ans
}

/**
 * @title 平衡集合数据结构
 * @description 用于高效维护可访问位置的集合
 */
type BalancedSet struct {
	data map[int]bool
}

func NewBalancedSet() *BalancedSet {
	return &BalancedSet{data: make(map[int]bool)}
}

func (s *BalancedSet) Insert(x int) {
	s.data[x] = true
}

func (s *BalancedSet) Remove(x int) {
	delete(s.data, x)
}

func (s *BalancedSet) Contains(x int) bool {
	return s.data[x]
}

/**
 * @description 辅助函数：返回两个数中的较大值
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * @description 辅助函数：返回两个数中的较小值
 */
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// 测试用例
	fmt.Println(minReverseOperations(4, 0, []int{1, 2}, 4)) // [0,-1,-1,1]
	fmt.Println(minReverseOperations(5, 1, []int{4}, 2))    // [1,0,1,2,-1]
	fmt.Println(minReverseOperations(5, 1, []int{4}, 3))    // [-1,0,-1,1,-1]
	fmt.Println(minReverseOperations(5, 1, []int{4}, 4))    // [-1,0,1,-1,-1]
	fmt.Println(minReverseOperations(5, 1, []int{4}, 5))    // [-1,0,-1,1,-1]
}
