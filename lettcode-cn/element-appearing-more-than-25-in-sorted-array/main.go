package main

import "fmt"

func findSpecialInteger(arr []int) int {
	result := arr[0]
	max := 1
	var cnt int = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			cnt++
		} else {
			if cnt > max {
				max = cnt
				result = arr[i-1]
			}
			cnt = 1
		}
	}
	if cnt > max {
		result = arr[len(arr)-1]
	}
	return result
}

func main() {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	fmt.Println(findSpecialInteger(arr))
}
