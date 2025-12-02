package main

import "testing"

func TestCountTrapezoids(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "示例1: 3个点在y=0, 2个点在y=2",
			points: [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}},
			want:   3,
		},
		{
			name:   "示例2: 2个点在y=0, 2个点在y=1",
			points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			want:   1,
		},
		{
			name:   "最小情况: 4个点组成1个梯形",
			points: [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
			want:   1,
		},
		{
			name:   "无法组成梯形: 每行只有1个点",
			points: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			want:   0,
		},
		{
			name:   "无法组成梯形: 只有一行有多个点",
			points: [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 1}},
			want:   0,
		},
		{
			name:   "三条水平线",
			points: [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}},
			// y=0: C(2,2)=1, y=1: C(2,2)=1, y=2: C(2,2)=1
			// 组合: 1*1 + 1*1 + 1*1 = 3 (y0-y1, y0-y2, y1-y2)
			want: 3,
		},
		{
			name:   "大量点在同一行",
			points: [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {0, 1}, {1, 1}},
			// y=0: C(4,2)=6, y=1: C(2,2)=1
			// 组合: 6 * 1 = 6
			want: 6,
		},
		{
			name:   "负坐标",
			points: [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}},
			want:   1,
		},
		{
			name:   "大坐标值",
			points: [][]int{{-100000000, 0}, {100000000, 0}, {-100000000, 100000000}, {100000000, 100000000}},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTrapezoids(tt.points)
			if got != tt.want {
				t.Errorf("countTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 基准测试
func BenchmarkCountTrapezoids(b *testing.B) {
	// 构造一个较大的测试用例: 1000个点，分布在10条水平线上
	points := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		points[i] = []int{i, i % 10}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countTrapezoids(points)
	}
}
