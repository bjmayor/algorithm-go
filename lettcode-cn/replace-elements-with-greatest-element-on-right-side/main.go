package main

import "fmt"

func replaceElements(arr []int) []int {
	return replaceElements2(arr)
}
func replaceElements1(arr []int) []int {
	s := []int{}
	for i := 1; i < len(arr); i++ {
		if len(s) == 0 {
			s = append(s, i)
		} else {
			for len(s) > 0 && arr[s[len(s)-1]] < arr[i] {
				s = s[:len(s)-1]
			}
			s = append(s, i)
		}

	}
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if len(s) == 0 {
			result[i] = -1
		} else {
			result[i] = arr[s[0]]
			if s[0] <= i+1 {
				s = s[1:]
			}
		}
	}
	return result

}

func replaceElements2(arr []int) []int {
	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		arr[i] = max
		if temp > max {
			max = temp
		}
	}
	return arr
}
func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}
