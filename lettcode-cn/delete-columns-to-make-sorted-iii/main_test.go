package main

import "testing"

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want int
	}{
		{
			name: "示例1",
			strs: []string{"babca", "bbazb"},
			want: 3,
		},
		{
			name: "示例2",
			strs: []string{"edcba"},
			want: 4,
		},
		{
			name: "示例3",
			strs: []string{"ghi", "def", "abc"},
			want: 0,
		},
		{
			name: "单列",
			strs: []string{"a", "b", "c"},
			want: 0,
		},
		{
			name: "全部需要删除",
			strs: []string{"zyx", "wvu", "tsr"},
			want: 2,
		},
		{
			name: "保留一列",
			strs: []string{"abb", "bcc"},
			want: 0,
		},
		{
			name: "已经排序",
			strs: []string{"abc", "bcd", "cde"},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
