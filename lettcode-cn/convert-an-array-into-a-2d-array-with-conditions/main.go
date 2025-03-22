package main

import "fmt"

/*
给你一个整数数组 nums 。请你创建一个满足以下条件的二维数组：

二维数组应该 只 包含数组 nums 中的元素。
二维数组中的每一行都包含 不同 的整数。
二维数组的行数应尽可能 少 。
返回结果数组。如果存在多种答案，则返回其中任何一种。

请注意，二维数组的每一行上可以存在不同数量的元素。
*/
func findMatrix(nums []int) [][]int {
	counts := make(map[int]int) // 创建一个 hashmap 来存储每个数字的出现次数
	for _, num := range nums {
		counts[num]++
	}

	ans := [][]int{}
	for len(counts) > 0 {
		row := []int{}
		for num := range counts {
			row = append(row, num) // 将当前的数字添加到行中
			counts[num]--           // 减少该数字的计数
			if counts[num] == 0 {
				delete(counts, num) // 如果计数为0，则从 hashmap 中删除该数字
			}
		}
		ans = append(ans, row) // 将行添加到结果中
	}
	return ans
}

func main() {
	nums := []int{1, 3, 4, 1, 2, 3, 1}
	fmt.Println(findMatrix(nums))
}
