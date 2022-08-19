package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	diffs := make([]int, n+2) // 差分数组
	for _, booking := range bookings {
		l, r, v := booking[0], booking[1], booking[2]
		diffs[l] += v
		diffs[r+1] -= v
	}
	ans := make([]int, n)
	pre := 0
	for i, diff := range diffs[1 : n+1] {
		ans[i] = pre + diff
		pre += diff
	}
	return ans
}

func main() {
	fmt.Println(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}
