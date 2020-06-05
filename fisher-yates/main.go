package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(indexes []int)  {
	for i:=len(indexes);i>0;i-- {
		lastIdx := i-1;
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

func main() {
	a := []int{3,5,8,11}
	rand.Seed(time.Now().UnixNano()) // 没有这个。rand.Intn 是固定的顺序。
	fmt.Println(a)
	shuffle(a)
	fmt.Println(a)
}
