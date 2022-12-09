package main

func largestAltitude(gain []int) (ans int) {
	ans = 0
	pre := 0
	for _, v := range gain {
		cur := v + pre
		pre = cur
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
