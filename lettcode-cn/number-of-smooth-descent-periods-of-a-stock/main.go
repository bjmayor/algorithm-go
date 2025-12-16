package main

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	if n == 0 {
		return 0
	}
	var ans int64 = 0
	size := 1
	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			size++
		} else {
			ans += int64(size * (size + 1) / 2)
			size = 1
		}
	}
	ans += int64(size * (size + 1) / 2)
	return ans
}

func main() {
	// 留空或用于手动调试
}
