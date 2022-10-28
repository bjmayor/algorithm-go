package main

func sumSubarrayMins(arr []int) int {
	left := make([]int, len(arr))
	right := make([]int, len(arr))

	stack := []int{}
	for i, v := range arr {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > v {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = len(arr)
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	var ans int
	for i, v := range arr {
		ans = (ans + v*(right[i]-i)*(i-left[i])) % (1e9 + 7)
	}
	return ans
}

func main() {
	println(sumSubarrayMins([]int{3, 1, 2, 4}))        // 17
	println(sumSubarrayMins([]int{11, 81, 94, 43, 3})) // 444
	println(sumSubarrayMins([]int{71, 55, 82, 55}))    //593
}
