package main

func totalFruit(fruits []int) int {
	m := map[int]int{}
	var ans int
	var left int
	for right, v := range fruits {
		m[v]++
		for len(m) > 2 {
			y := fruits[left]
			m[y]--
			if m[y] == 0 {
				delete(m, y)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	println(totalFruit([]int{1, 2, 1}))
	println(totalFruit([]int{0, 1, 2, 2}))
	println(totalFruit([]int{1, 2, 3, 2, 2}))
	println(totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
