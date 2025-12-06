package main

import "testing"

func TestCountPartitions(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		// 示例测试用例
		{
			name:     "Example 1",
			nums:     []int{9, 4, 1, 3, 7},
			k:        4,
			expected: 6,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 3, 4},
			k:        0,
			expected: 2,
		},

		// 边界情况
		{
			name:     "Minimum array length",
			nums:     []int{1, 2},
			k:        1,
			expected: 2,
		},
		{
			name:     "All elements the same",
			nums:     []int{5, 5, 5, 5},
			k:        0,
			expected: 8, // 2^(n-1) = 2^3 = 8，每个位置都可以分割或不分割
		},
		{
			name:     "Large k, all in one partition",
			nums:     []int{1, 2, 3, 4, 5},
			k:        100,
			expected: 16, // 所有数都可以在一个分割中，有多种分割方式
		},
		{
			name:     "k = 0, all different",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: 1, // 每个元素必须单独分割
		},
		{
			name:     "k = 0, some same",
			nums:     []int{1, 1, 2, 2},
			k:        0,
			expected: 4, // [1],[1],[2],[2] | [1,1],[2],[2] | [1],[1],[2,2] | [1,1],[2,2]
		},

		// 中等复杂度
		{
			name:     "Mixed values",
			nums:     []int{1, 5, 2, 4, 3},
			k:        2,
			expected: 4,
		},
		{
			name:     "Two elements, extreme difference",
			nums:     []int{1, 100},
			k:        10,
			expected: 1, // 极差=99>10，所以只能 [1],[100]
		},

		// 递增和递减序列
		{
			name:     "Increasing sequence",
			nums:     []int{1, 2, 3, 4},
			k:        1,
			expected: 5, // 相邻元素极差都是1，可以有多种分割方式
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{4, 3, 2, 1},
			k:        1,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countPartitions(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("countPartitions(%v, %d) = %d, want %d",
					tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

// 性能测试
func BenchmarkCountPartitions(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = (i % 100) + 1
	}
	k := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countPartitions(nums, k)
	}
}

// 边界值测试
func TestCountPartitionsEdgeCases(t *testing.T) {
	t.Run("Single pair boundary", func(t *testing.T) {
		// nums[i..j] 的极差刚好等于 k
		nums := []int{1, 6}
		k := 5
		result := countPartitions(nums, k)
		expected := 2 // [1],[6] 或 [1,6]
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("Single pair boundary exceeded", func(t *testing.T) {
		// nums[i..j] 的极差超过 k
		nums := []int{1, 6}
		k := 4
		result := countPartitions(nums, k)
		expected := 1 // 只有 [1],[6]
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})

	t.Run("Large values", func(t *testing.T) {
		nums := []int{1000000000, 1000000001}
		k := 1
		result := countPartitions(nums, k)
		expected := 2 // [1000000000],[1000000001] 或 [1000000000,1000000001]
		if result != expected {
			t.Errorf("Expected %d, got %d", expected, result)
		}
	})
}

// 验证 MOD 操作
func TestCountPartitionsMOD(t *testing.T) {
	t.Run("Result with modulo", func(t *testing.T) {
		// 创建一个会产生大数字的测试案例
		nums := make([]int, 20)
		for i := 0; i < 20; i++ {
			nums[i] = i + 1
		}
		k := 1

		result := countPartitions(nums, k)
		// 结果应该小于 MOD (10^9 + 7)
		MODValue := 1000000007
		if result < 0 || result >= MODValue {
			t.Errorf("Result %d should be in [0, %d)", result, MODValue)
		}
	})
}
