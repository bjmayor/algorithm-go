package main

import "fmt"
//栈做法
//func longestValidParentheses(s string) int {
//	var stack []int
//	var res int
//	stack = append(stack, -1)
//	for i:=0;i<len(s);i++ {
//		if s[i] == '('  {
//			stack = append(stack, i)
//		} else {
//			stack = stack[:len(stack)-1] //弹出'(', 配对
//			if len(stack) == 0 {//
//				stack = append(stack, i)//有效配对的前一个index
//			} else {
//				if i-stack[len(stack)-1]	> res {
//					res = i-stack[len(stack)-1]
//				}
//			}
//		}
//	}
//	return res
//}

//双指针做法
func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2 * right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2 * left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func main()  {
	tests := []struct{
		S string
		Expected int
	} {
		{ ")()(",2},
		{ "()(()",2},
		{ "(()",2},
		{ ")()())",4},
		{ ")()())()()(",4},
	}
	for i, t := range tests {
		var real = longestValidParentheses(t.S)
		if real !=  t.Expected {
			fmt.Println(i," expected:", t.Expected, " real:", real, t)
		}
	}
}