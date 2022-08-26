package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - 1
	toremove := len(arr) - k
	for i := 1; i <= toremove; i++ {
		if abs(arr[right]-x) >= abs(arr[left]-x) {
			right--
		} else {
			left++
		}
	}
	return arr[left : right+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}
