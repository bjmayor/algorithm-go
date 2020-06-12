package main

import (
	"fmt"
)

//单调栈 解法， a 维护单调递减
func dailyTemperatures(T []int) []int {
	result := make([]int, len(T), len(T)) //存结果
	stack := make([]int, 0, len(T)) //存索引
	for i:=0;i<len(T);i++ {
		top := len(stack)-1
		if len(stack)==0 || T[i] <= T[stack[top]] {
			stack = append(stack, i)
		} else  {
			for top >=0 && T[i] > T[stack[top]]{
				result[stack[top]] = i-stack[top]
				stack = stack[:top]
				top--
			}
			stack = append(stack, i)
		}
	}
	for _,i := range stack {
		result[i] = 0
	}
	return result
}

//从后往前推
func dailyTemperatures2(T []int) []int {
	result := make([]int, len(T), len(T)) //存结果
	result[len(T)-1] = 0
	for i:=len(T)-2;i>=0;i-- {
		if T[i] < T[i+1] {
			result[i] = 1
		} else {
			j := i+1  // 第一个比T[i]大的
			for j <len(result){
				if T[j] > T[i] {
					result[i] = j-i
					break
				} else {
					//没有更大的，退出
					if result[j] == 0 {
						break
					}
					j += result[j]
				}
			}
		}
	}
	return result
}

func main()  {
	fmt.Println(dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}