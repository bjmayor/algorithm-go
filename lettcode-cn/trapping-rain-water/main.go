package main

import "fmt"

/**
 * trap 函数用于计算给定高度数组中可以储存的雨水总量。
 * 算法思想：使用不同的算法（如栈、双指针等）来计算雨水量。
 *
 * @param height 整数数组，表示每个位置的高度
 * @return int 返回可以储存的雨水总量
 */
func trap(height []int) int {
	return trap4(height)
}

/**
 * trap1 使用栈来计算雨水量。
 * 算法思想：通过栈来维护一个递减序列，当遇到比栈顶大的元素时，计算可以储存的雨水量。
 *
 * @param height 整数数组，表示每个位置的高度
 * @return int 返回可以储存的雨水总量
 */
func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	s := make([]int, 0)
	var res int
	for i := 0; i < len(height); i++ {
		if len(s) == 0 {
			s = append(s, height[i])
		} else {
			top := s[len(s)-1]
			if height[i] >= top {
				// 不能再积更多雨水了。开始计算, 左低右高
				if height[i] >= s[0] {
					v := s[0]
					for len(s) > 0 {
						top := s[len(s)-1]
						s = s[:len(s)-1]
						if v > top {
							res += v - top
						}
					}
					s = append(s, height[i])
				} else {
					// 连续上升的也入栈
					s = append(s, height[i])
				}
			} else {
				// 下降的直接入栈
				s = append(s, height[i])
			}
		}
	}
	if len(s) > 0 {
		for i := 0; i < len(s)/2; i++ {
			s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
		}
		return res + trap(s)
	}
	return res
}

/**
 * trap2 使用栈来计算雨水量，算法更简洁。
 * 算法思想：通过栈来维护一个递减序列，计算可以储存的雨水量。
 *
 * @param height 整数数组，表示每个位置的高度
 * @return int 返回可以储存的雨水总量
 */
func trap2(height []int) int {
	s := make([]int, 0) // 存索引，都是下降序列
	var res int
	i := 0
	for i < len(height) {
		for len(s) > 0 && height[i] > height[s[len(s)-1]] {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 {
				break
			}
			pre := s[len(s)-1]
			distance := i - pre - 1
			min := height[i]
			if height[pre] < min {
				min = height[pre]
			}
			res += distance * (min - height[top])
		}
		s = append(s, i)
		i++
	}
	return res
}

/**
 * trap3 使用双指针来计算雨水量。
 * 算法思想：通过左右指针来维护最大高度，计算可以储存的雨水量。
 *
 * @param height 整数数组，表示每个位置的高度
 * @return int 返回可以储存的雨水总量
 */
func trap3(height []int) int {
	left := 0
	right := len(height) - 1
	var leftMax, rightMax int
	var ans int
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}

/**
 * trap4 使用降序再升序的方式来计算雨水量。
 * 算法思想：通过遍历高度数组，计算可以储存的雨水量。
 *
 * @param height 整数数组，表示每个位置的高度
 * @return int 返回可以储存的雨水总量
 */
func trap4(height []int) int {
	water := 0
	temp := 0
	left := 0
	// 降序 再 升序，中间的就是雨水
	for i, h := range height {
		if height[left] == 0 {
			left = i
			temp += height[left] - h
		} else {
			if h >= height[left] {
				water += temp
				left = i
				temp = 0
			} else {
				temp += height[left] - h
			}
		}
	}
	if temp > 0 {
		temp = 0
		right := len(height) - 1
		for i := right; i >= 0; i-- {
			h := height[i]
			if height[right] == 0 {
				right = i
				temp += height[right] - h
			} else {
				if h > height[right] {
					water += temp
					right = i
					temp = 0
				} else {
					temp += height[right] - h
				}
			}
		}
	}
	return water
}

func main() {
	fmt.Println(trap([]int{
		0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
	})) // 6

	fmt.Println(trap([]int{
		5, 2, 1, 2, 1, 5,
	})) // 14
	fmt.Println(trap([]int{
		5, 4, 1, 2,
	})) // 1
	fmt.Println(trap([]int{
		2, 3, 4, 1, 6, 6, 6, 1, 0, 1,
	})) // 4
	fmt.Println(trap([]int{
		5, 5, 1, 7, 1, 1, 5, 2, 7, 6,
	})) // 23
}
