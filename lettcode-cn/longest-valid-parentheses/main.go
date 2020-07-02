package main

import "fmt"

func longestValidParentheses(s string) int {
	var stack []int
	var res int
	for i:=0;i<len(s);i++ {
		if s[i] == '(' {
			stack = append(stack,'(')
		} else {
			if len(stack) > 0 {
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
					stack = append(stack, 2+')')
				} else {
					cnt := 0
					found := false
					for len(stack) > 0 {
						top :=stack[len(stack)-1]
						if top != '(' {
							cnt += top-')'
							stack = stack[:len(stack)-1]
						} else {
							found = true
							cnt += 2
							stack = stack[:len(stack)-1]
							break
						}
					}
					if found {
						stack = append(stack, cnt+')')
					} else {
						if cnt > res {
							res = cnt
						}
					}
				}
			}
		}
	}
	cnt := 0
	for len(stack) > 0 {
		top :=stack[len(stack)-1]
		if top != '(' {
			cnt += top-')'
			stack = stack[:len(stack)-1]
			if cnt > res {
				res = cnt
			}
		} else {
			stack = stack[:len(stack)-1]
			cnt = 0
		}
	}
	return res
}

func main()  {
	tests := []struct{
		S string
		Expected int
	} {
		//{ "()(()",2},
		//{ "(()",2},
		//{ ")()())",4},
		{ ")()())()()(",4},
	}
	for i, t := range tests {
		var real = longestValidParentheses(t.S)
		if real !=  t.Expected {
			fmt.Println(i," expected:", t.Expected, " real:", real, t)
		}
	}
}