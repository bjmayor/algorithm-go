package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, v := range asteroids {
		if v > 0 {
			stack = append(stack, v)
		} else {
			ntop := v
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				if top > 0 {
					stack = stack[:len(stack)-1]
					if abs(top) < abs(v) {
						ntop = v
						continue
					} else if abs(top) == abs(v) {
						ntop = 0
						break
					} else {
						ntop = top
						break
					}
				} else {
					break
				}
			}
			if ntop != 0 {
				stack = append(stack, ntop)
			}
		}
	}
	return stack
}

func abs(x int) int  {
	if x > 0 {
		return x
	}
	return -x
}
func main() {
	//fmt.Println(asteroidCollision([]int{5, 10, -5})) // [5,10]
	//fmt.Println(asteroidCollision([]int{8, -8}))     // []
	//fmt.Println(asteroidCollision([]int{10, 2, -5})) // [10]
	fmt.Println(asteroidCollision([]int{-2,-1,1,2})) // [[-2,-1,1,2]]
}
