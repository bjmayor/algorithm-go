package main

import "fmt"

func divingBoard(shorter int, longer int, k int) []int {
	if k<=0 {
		return  nil
	}
	if shorter == longer {
		return []int{shorter*k}
	}
	var result []int
	for i:=k;i>=0;i-- {
		result = append(result, i*shorter+(k-i)*longer)
	}
	return result
}

func main()  {
	fmt.Println(divingBoard(1,2,3))
	fmt.Println(divingBoard(1,1,10000))
}