package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

func main()  {
	fmt.Println(isPowerOfTwo(18))
	fmt.Println(isPowerOfTwo(16))
}