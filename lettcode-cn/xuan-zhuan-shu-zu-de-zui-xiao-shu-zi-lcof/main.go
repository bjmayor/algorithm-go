package main

import "fmt"

func minArray(numbers []int) int {
	l := 0
	r := len(numbers)-1
	for l < r {
		mid := (l+r)/2 //假设是这个值
		if numbers[mid] > numbers[r] {
			l = mid+1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] == numbers[r] {
			r--
		}
	}
	return numbers[l]
}

func main()  {
	fmt.Println(minArray([]int{3,4,5,1,2}))//1
	fmt.Println(minArray([]int{2,2,2,0,1}))//0
	fmt.Println(minArray([]int{3,3,1,3}))//1
}
