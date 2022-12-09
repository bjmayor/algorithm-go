package main

const mod int = 1e9 + 7

func nthMagicalNumber(n int, a int, b int) int {
	l := min(a, b)
	r := n * min(a, b)
	c := a / gcd(a, b) * b // 最小公倍数
	for l <= r {
		mid := (l + r) / 2
		cnt := mid/a + mid/b - mid/c // <=mid 的神奇数字个数
		if cnt >= n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return (r + 1) % mod

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 最大公因子
func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
