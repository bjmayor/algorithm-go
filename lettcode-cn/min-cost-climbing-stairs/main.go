package main

func minCostClimbingStairs(cost []int) int {
	p, q := 0, 0
	for i := 2; i <= len(cost); i++ {
		p, q = q, min(cost[i-2]+p, cost[i-1]+q)
	}
	return q
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	//println(minCostClimbingStairs([]int{10, 15, 20})) //15
	println(minCostClimbingStairs([]int{1, 100})) //1
}
