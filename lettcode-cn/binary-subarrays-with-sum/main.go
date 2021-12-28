package main

// 哈希法
func hash(nums []int, goal int) int {
	m := make(map[int]int)

	total := 0
	result := 0
	for _,v := range nums {
		total += v
		if total == goal {
			result+=1
		}
		delta := total-goal
		if v, ok := m[delta]; ok {
			result += v
		}
		m[total] = m[total]+1
	}

	return result
}

func slideWindow(nums []int, goal int) (ans int) {
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0 // sum1 为 left1->right 的和, 其和<=goal, sum2 为left2->right 的和，其和<goal , 那么left1->left2 区间的数都满足。
	for right, num := range nums {
		sum1 += num
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}
		sum2 += num
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return
}

func numSubarraysWithSum(nums []int, goal int) (ans int) {
	return slideWindow(nums, goal)
}
func main()  {
	println(numSubarraysWithSum([]int{1,0,1,0,1},2))
	println(numSubarraysWithSum([]int{0,0,0,0,0},0))
}

