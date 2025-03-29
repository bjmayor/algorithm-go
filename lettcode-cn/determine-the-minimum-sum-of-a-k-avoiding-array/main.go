package main

import "fmt"

/*
2829. k-avoiding 数组的最小总和.
给你两个整数 n 和 k 。

对于一个由 不同 正整数组成的数组，如果其中不存在任何求和等于 k 的不同元素对，则称其为 k-avoiding 数组。

返回长度为 n 的 k-avoiding 数组的可能的最小总和。
*/
func minimumSum(n int, k int) int {
	sum := 0
	added := make(map[int]bool)
	cur := 1
	total := 0
	for total < n {
		if cur < k && added[k-cur] {
			cur++
			continue
		}
		added[cur] = true
		sum += cur
		total++
		cur++
	}
	return sum
}

func main() {
	fmt.Println(minimumSum(5, 4)) // 18
	fmt.Println(minimumSum(2, 6)) //3
}
