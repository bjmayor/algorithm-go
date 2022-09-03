package main

import "fmt"

func permute(nums []int) [][]int {
	ans := [][]int{{}}
	for _, v := range nums {
		tmp := [][]int{}
		for _, item := range ans {
			for j := 0; j < len(item)+1; j++ {
				add := make([]int, len(item)+1)
				copy(add, item[:j])
				add[j] = v
				copy(add[j+1:], item[j:])
				tmp = append(tmp, add)
			}
		}
		ans = tmp
	}
	return ans
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
