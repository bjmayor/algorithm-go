package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	totalGas, totalCost, currentGas, start := 0, 0, 0, 0

	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]

		// 如果当前油量小于 0，更新起始点
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	// 如果总油量小于总消耗量，返回 -1
	if totalGas < totalCost {
		return -1
	}
	return start

}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
