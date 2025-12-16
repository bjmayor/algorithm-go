package main

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name     string
		corridor string
		expected int
	}{
		{
			name:     "示例1: SSPPSPS",
			corridor: "SSPPSPS",
			expected: 3,
		},
		{
			name:     "示例2: PPSPSP",
			corridor: "PPSPSP",
			expected: 1,
		},
		{
			name:     "示例3: S",
			corridor: "S",
			expected: 0,
		},
		{
			name:     "没有座位",
			corridor: "PPPP",
			expected: 0,
		},
		{
			name:     "两个座位中间有两个植物",
			corridor: "SPPSP",
			expected: 1,
		},
		{
			name:     "三个座位-奇数",
			corridor: "SPSPS",
			expected: 0,
		},
		{
			name:     "恰好两个座位",
			corridor: "SS",
			expected: 1,
		},
		{
			name:     "两个座位中间有植物",
			corridor: "SPPPS",
			expected: 1,
		},
		{
			name:     "多段情况",
			corridor: "SSPPSSPPSS",
			expected: 9, // 第一段SS和第二段SS之间2个P(3种)，第二段SS和第三段SS之间2个P(3种)，3*3=9
		},
		{
			name:     "多段中间有多个植物",
			corridor: "SSPPPPSSPPPPSS",
			expected: 25, // (4+1) * (4+1) = 25
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberOfWays(tt.corridor)
			if result != tt.expected {
				t.Errorf("numberOfWays(%q) = %d, expected %d", tt.corridor, result, tt.expected)
			}
		})
	}
}
