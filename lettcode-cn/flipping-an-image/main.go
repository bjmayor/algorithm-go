package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i, arr := range A {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = 1-arr[len(arr)-1-i], 1-arr[i]
		}
		if len(arr)%2 == 1 {
			arr[len(arr)/2] = 1 - arr[len(arr)/2]
		}
		A[i] = arr
	}
	return A
}

func main() {
	fmt.Print(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}
