package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAllPeople(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		meetings    [][]int
		firstPerson int
		want        []int
	}{
		{
			name:        "示例1",
			n:           6,
			meetings:    [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 5},
		},
		{
			name:        "示例2",
			n:           4,
			meetings:    [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}},
			firstPerson: 3,
			want:        []int{0, 1, 3},
		},
		{
			name:        "示例3",
			n:           5,
			meetings:    [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 4},
		},
		{
			name:        "只有初始两人",
			n:           3,
			meetings:    [][]int{{1, 2, 1}},
			firstPerson: 2,
			want:        []int{0, 1, 2},
		},
		{
			name:        "同一时间多个会议",
			n:           6,
			meetings:    [][]int{{1, 2, 1}, {2, 3, 1}, {3, 4, 1}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 4},
		},
		{
			name:        "不连通的会议",
			n:           5,
			meetings:    [][]int{{3, 4, 2}, {1, 2, 1}},
			firstPerson: 2,
			want:        []int{0, 1, 2},
		},
		{
			name:        "时间先后导致传播",
			n:           6,
			meetings:    [][]int{{0, 2, 1}, {1, 3, 1}, {4, 5, 1}, {2, 3, 2}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3},
		},
		{
			name:        "复杂的时间序列",
			n:           11,
			meetings:    [][]int{{5, 1, 4}, {2, 4, 5}, {3, 7, 6}, {8, 10, 7}, {6, 8, 7}, {3, 5, 1}, {5, 2, 7}, {2, 5, 7}},
			firstPerson: 9,
			want:        []int{0, 9},
		},
		{
			name:        "最小情况",
			n:           2,
			meetings:    [][]int{},
			firstPerson: 1,
			want:        []int{0, 1},
		},
		{
			name:        "循环传播",
			n:           6,
			meetings:    [][]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}, {4, 5, 4}},
			firstPerson: 1,
			want:        []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllPeople(tt.n, tt.meetings, tt.firstPerson)
			// 排序结果以便比较
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestUnionFind 测试并查集基本操作
func TestUnionFind(t *testing.T) {
	t.Run("基本操作", func(t *testing.T) {
		uf := NewUnionFind(5)

		// 初始时每个节点独立
		for i := 0; i < 5; i++ {
			if uf.Find(i) != i {
				t.Errorf("初始时节点 %d 的根应该是自己", i)
			}
		}

		// 合并 0 和 1
		uf.Union(0, 1)
		if !uf.IsConnected(0, 1) {
			t.Error("0 和 1 应该连通")
		}

		// 合并 2 和 3
		uf.Union(2, 3)
		if !uf.IsConnected(2, 3) {
			t.Error("2 和 3 应该连通")
		}

		// 0 和 2 不应该连通
		if uf.IsConnected(0, 2) {
			t.Error("0 和 2 不应该连通")
		}

		// 合并 1 和 3，此时 0,1,2,3 都连通
		uf.Union(1, 3)
		if !uf.IsConnected(0, 3) {
			t.Error("0 和 3 应该连通")
		}
		if !uf.IsConnected(1, 2) {
			t.Error("1 和 2 应该连通")
		}
	})

	t.Run("重置操作", func(t *testing.T) {
		uf := NewUnionFind(5)

		// 合并 0, 1, 2
		uf.Union(0, 1)
		uf.Union(1, 2)

		if !uf.IsConnected(0, 2) {
			t.Error("0 和 2 应该连通")
		}

		// 重置节点 2
		uf.Reset(2)

		// 2 应该与 0 和 1 不再连通
		if uf.IsConnected(0, 2) {
			t.Error("重置后 0 和 2 不应该连通")
		}
		if uf.IsConnected(1, 2) {
			t.Error("重置后 1 和 2 不应该连通")
		}

		// 但 0 和 1 仍然连通
		if !uf.IsConnected(0, 1) {
			t.Error("0 和 1 应该仍然连通")
		}
	})
}

// BenchmarkFindAllPeople 性能测试
func BenchmarkFindAllPeople(b *testing.B) {
	n := 1000
	meetings := make([][]int, 0)
	for i := 1; i < n; i++ {
		meetings = append(meetings, []int{i - 1, i, i})
	}
	firstPerson := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findAllPeople(n, meetings, firstPerson)
	}
}
