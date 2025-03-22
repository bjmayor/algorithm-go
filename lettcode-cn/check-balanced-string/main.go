package main

import "fmt"

func isBalanced(num string) bool {
	odd, even := 0, 0

	for i := 0; i < len(num); i++ {
		if i%2 == 0 {
			even += int(num[i] - '0')
		} else {
			odd += int(num[i] - '0')
		}
	}
	return odd == even
}

func main() {
	fmt.Println(isBalanced("1234"))  //false
	fmt.Println(isBalanced("24123")) //true
}
