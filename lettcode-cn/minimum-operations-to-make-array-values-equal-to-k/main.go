package main

import "fmt"

func minOperations(nums []int, k int) int {
	biggerNums := 0
	seen := make(map[int]bool)
	for _, num := range nums {
		if num < k {
			return -1
		}
		if _, ok := seen[num]; !ok {
			seen[num] = true
			if num > k {
				biggerNums += 1
			}
		}
	}

	return biggerNums
}

func main() {
	fmt.Println(minOperations([]int{5, 2, 5, 4, 5}, 2)) //2
	fmt.Println(minOperations([]int{2, 1, 2}, 2))       //-1
	fmt.Println(minOperations([]int{9, 7, 5, 3}, 1))    //4
}
