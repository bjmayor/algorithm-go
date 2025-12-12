package main

import (
	"reflect"
	"testing"
)

func TestCountMentions(t *testing.T) {
	tests := []struct {
		name           string
		numberOfUsers  int
		events         [][]string
		expectedResult []int
	}{
		{
			name:          "基本功能测试 - 提及所有用户",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "100", "ALL"},
			},
			expectedResult: []int{1, 1, 1},
		},
		{
			name:          "提及在线用户",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "100", "HERE"},
			},
			expectedResult: []int{1, 1, 1},
		},
		{
			name:          "提及特定用户",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "100", "id0"},
				{"MESSAGE", "101", "id2"},
			},
			expectedResult: []int{1, 0, 1},
		},
		{
			name:          "混合提及方式",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "100", "id0"},
				{"MESSAGE", "101", "ALL"},
				{"MESSAGE", "102", "HERE"},
			},
			expectedResult: []int{3, 2, 2},
		},
		{
			name:          "用户离线测试",
			numberOfUsers: 3,
			events: [][]string{
				{"OFFLINE", "50", "1"}, // 用户1离线
				{"MESSAGE", "100", "ALL"},
				{"MESSAGE", "110", "HERE"}, // 用户1重新上线
			},
			expectedResult: []int{2, 2, 2}, // 用户1在时间110重新上线
		},
		{
			name:          "用户重新上线测试",
			numberOfUsers: 3,
			events: [][]string{
				{"OFFLINE", "0", "1"}, // 用户1离线
				{"MESSAGE", "10", "ALL"},
				{"MESSAGE", "50", "HERE"}, // 用户1仍然离线
				{"MESSAGE", "70", "HERE"}, // 用户1重新上线
			},
			expectedResult: []int{3, 2, 3},
		},
		{
			name:          "多个用户同时离线",
			numberOfUsers: 4,
			events: [][]string{
				{"OFFLINE", "0", "0"},
				{"OFFLINE", "0", "2"},
				{"MESSAGE", "10", "ALL"},
				{"MESSAGE", "30", "HERE"},
				{"MESSAGE", "70", "HERE"}, // 用户0和2应该重新上线
			},
			expectedResult: []int{2, 3, 2, 3},
		},
		{
			name:           "空事件列表",
			numberOfUsers:  2,
			events:         [][]string{},
			expectedResult: []int{0, 0},
		},
		{
			name:          "无提及的消息",
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "100", ""},
			},
			expectedResult: []int{0, 0},
		},
		{
			name:          "时间顺序测试",
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "200", "ALL"},
				{"MESSAGE", "100", "ALL"},
			},
			expectedResult: []int{2, 2},
		},
		{
			name:          "一条消息提及多个用户",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "10", "id1 id0"},
			},
			expectedResult: []int{1, 1, 0},
		},
		{
			name:          "在线离线时间边界测试",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "2", "HERE"},
				{"OFFLINE", "2", "1"},
				{"OFFLINE", "1", "0"},
				{"MESSAGE", "61", "HERE"},
			},
			expectedResult: []int{1, 0, 2},
		},
		{
			name:          "case287",
			numberOfUsers: 3,
			events: [][]string{
				{"MESSAGE", "5", "HERE"},
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "15", "HERE"},
				{"OFFLINE", "18", "2"},
				{"MESSAGE", "20", "HERE"},
			},
			expectedResult: []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countMentions(tt.numberOfUsers, tt.events)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("countMentions() = %v, expected %v", result, tt.expectedResult)
			}
		})
	}
}
