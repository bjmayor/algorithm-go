package main

func validateStackSequences(pushed []int, popped []int) bool {
	j := 0
	stack := []int{}
	for _, x := range pushed {
		stack = append(stack, x)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

func main() {
	println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})) // true
	println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})) // false
	println(validateStackSequences([]int{1, 0}, []int{1, 0}))                   // false
}
