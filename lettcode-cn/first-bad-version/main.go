package main

import "fmt"

func isBadVersion(version int) bool {
	if version >= 1 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	//定义left-right 为都是正确的版本。
	for right>=left {
		mid := (right+left)/2
		if isBadVersion(mid) {
			right = mid-1
		} else {
			left = mid+1
		}
	}
	return right+1
}

func main()  {
	fmt.Println(firstBadVersion(1))
}