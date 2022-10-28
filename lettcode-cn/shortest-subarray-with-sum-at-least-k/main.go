package main

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSumArr := make([]int, n+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
	}
	ans := n + 1
	q := []int{}
	for i, curSum := range preSumArr {
		// q 中元素做为起点
		for len(q) > 0 && curSum-preSumArr[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:] // 后面就算满足需求，长度也只会更大，所以可以直接干掉
		}
		// 如果q中元素比curSum 还要大，可以排除掉, 保持单调递增
		for len(q) > 0 && preSumArr[q[len(q)-1]] >= curSum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans < n+1 {
		return ans
	}
	return -1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//println(shortestSubarray([]int{1}, 1))                     //1
	//println(shortestSubarray([]int{1, 2}, 4))                  //-1
	//println(shortestSubarray([]int{2, -1, 2}, 3))              //3
	//println(shortestSubarray([]int{77, 19, 35, 10, -14}, 19))  //1
	println(shortestSubarray([]int{84, -37, 32, 40, 95}, 167)) //3
}
