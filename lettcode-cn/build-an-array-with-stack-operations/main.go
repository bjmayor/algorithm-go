package main

import "fmt"

func buildArray(target []int, n int) (ans []string) {
	cur := 1
	for _, v := range target {
		for cur < v {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
			cur++
		}
		ans = append(ans, "Push")
		cur++
	}
	return
}

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}
