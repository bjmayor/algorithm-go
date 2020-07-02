package main

import (
	"fmt"
	"math"
)

func sumNums(n int) int {
	return 	(int(math.Pow(float64(n),2))+n) >> 1
}

func main()  {
	fmt.Println(sumNums(100))
}
