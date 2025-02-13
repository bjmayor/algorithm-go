package main

import "fmt"

/**
 * numOfMinutes 函数计算从头节点开始通知所有员工所需的时间。
 * @param n 员工的数量
 * @param headID 头节点的ID
 * @param manager 每个员工的直接经理的ID
 * @param informTime 每个员工通知其下属所需的时间
 * @return 返回通知所有员工所需的总时间
 */
func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	// 构建员工的下属列表
	subordinates := make([][]int, n)
	for i, m := range manager {
		if m != -1 {
			subordinates[m] = append(subordinates[m], i)
		}
	}

	// 使用队列进行广度优先搜索
	type state struct {
		empID int
		time  int
	}
	queue := []state{{headID, 0}}
	maxTime := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// 更新最大时间
		if current.time > maxTime {
			maxTime = current.time
		}

		// 将当前员工的下属加入队列
		for _, subID := range subordinates[current.empID] {
			queue = append(queue, state{subID, current.time + informTime[current.empID]})
		}
	}

	return maxTime
}

func main() {
	fmt.Println(numOfMinutes(11, 4, []int{5, 9, 6, 10, -1, 8, 9, 1, 9, 3, 4}, []int{0, 213, 0, 253, 686, 170, 975, 0, 261, 309, 337}))
}
