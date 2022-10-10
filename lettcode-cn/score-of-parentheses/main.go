package main

func scoreOfParentheses(s string) int {
	stack := []int{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, '(')
		} else {
			if stack[len(stack)-1] == '(' {
				v := 1
				stack = stack[:len(stack)-1]
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					top := stack[len(stack)-1] - '0'
					v += top
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, v+'0')
			} else {
				v := stack[len(stack)-1] - '0'
				v *= 2
				stack = stack[:len(stack)-2]
				for len(stack) > 0 && stack[len(stack)-1] != '(' {
					top := stack[len(stack)-1] - '0'
					v += top
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, v+'0')
			}
		}

	}
	return stack[0] - '0'
}

func main() {
	println(scoreOfParentheses("(()(()))"))
}
