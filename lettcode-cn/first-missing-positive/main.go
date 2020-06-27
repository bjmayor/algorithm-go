package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			//fmt.Println(num-1)
			nums[num - 1] = -abs(nums[num - 1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/* 这个用到了辅助空间过滤重复的数
func firstMissingPositive(nums []int) int {
	//先干掉负数 和 重复的数
	j := 0 //第一个非负数, 负数(重复的数).... 非负数...
	m := make(map[int]bool)
	for i:=0;i<len(nums);i++ {
		if nums[i]<=0  || m[nums[i]]{
			nums[i], nums[j] = nums[j], nums[i]
			j++
		} else {
			m[nums[i]] = true
		}
	}
	return firstMissingPositiveHelper(nums[j:], 1)
}
//从不重复的非负整数里找最小正数
func firstMissingPositiveHelper(nums []int, start int) int  {
	if len(nums) == 0 {
		return start
	}
	if len(nums) == 1 {
		if nums[0] != start  {
			return start
		} else {
			return start+1
		}
	}
	k := nums[0]
	i:=1//i之前的都<k
	j:=len(nums)-1 //j 之后的都是>=k
	for i<=j {
		if nums[i] > k {
			nums[i], nums[j] = nums[j],nums[i]
			j--
		} else {
			i++
		}
		if i>j {
			break
		}
		if nums[j] < k {
			nums[i], nums[j] = nums[j],nums[i]
			i++
		} else {
			j--
		}
	}
	nums[0], nums[i-1] = nums[i-1], nums[0]

	if k-start+1 == i {
		return firstMissingPositiveHelper(nums[i:], k+1)
	} else {
		return firstMissingPositiveHelper(nums[:i], start)
	}
}
 */

func main()  {
	fmt.Println(firstMissingPositive([]int{1,2,2,2,2,0}))//3
	fmt.Println(firstMissingPositive([]int{3,4,-1,1}))//2
	fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))//1
}
