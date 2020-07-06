package main

import "fmt"

//栈算法1
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	s := make([]int, 0)
	var res int
	for i:=0;i<len(height);i++ {
		if len(s) == 0 {
			s = append(s, height[i])
		} else {
			top := s[len(s)-1]
			if height[i] >= top {
				//不能再积更多雨水了。开始计算, 左低右高
				if height[i] >= s[0] {
					v := s[0]
					for len(s) > 0 {
						top := s[len(s)-1]
						s = s[:len(s)-1]
						if v > top {
							res += v-top
						}
					}
					s = append(s, height[i])

				} else {
					//连续上升的也入栈
					s= append(s, height[i])
				}
			} else {
				//下降的直接入栈
				s = append(s, height[i])
			}
		}
	}
	if len(s) > 0 {
		for i:=0;i<len(s)/2;i++ {
			s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
		}
		return res+trap(s)
	}
	return res
}

//栈算法2， 更简洁，但是效率没前一个高。主要是多了乘法
func trap2(height []int) int {
	s := make([]int, 0) //存索引，都是下降序列
	var res int
	i:=0
	for i<len(height){
		for len(s) >0 && height[i] > height[s[len(s)-1]]{
			top := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 {
				break
			}
			pre := s[len(s)-1]
			distance := i-pre-1
			min := height[i]
			if height[pre] < min {
				min = height[pre]
			}
			res += distance * (min-height[top])
		}
		s = append(s, i)
		i++
	}
	return res
}

//双指针
func trap3(height []int) int {
	left := 0
	right := len(height)-1
	var leftMax, rightMax int
	var ans int
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax-height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax-height[right]
			}
			right--
		}
	}
	return ans
}
func main()  {
	fmt.Println(trap([]int{
		0,1,0,2,1,0,1,3,2,1,2,1,
	}))//6

	fmt.Println(trap([]int{
		5,2,1,2,1,5,
	}))//14
	fmt.Println(trap([]int{
		5,4,1,2,
	}))//1
	fmt.Println(trap([]int{
		2,3,4,1,6,6,6,1,0,1,
	}))//4
}
