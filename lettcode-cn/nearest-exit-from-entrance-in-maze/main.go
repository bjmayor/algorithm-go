package main

import "fmt"

/**
https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/?envType=study-plan-v2&envId=graph-theory
1926. 迷宫中离入口最近的出口

给定一个 m x n 的迷宫，迷宫中有一个入口和一个出口。

迷宫中每个位置可以是：
- '.' 表示可以通行的路径
- '+' 表示障碍物
**/

func nearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	cols := len(maze[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}
	queue := [][]int{}
	queue = append(queue, []int{entrance[0], entrance[1], 0})
	visited[entrance[0]][entrance[1]] = true
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右，下，左，上
	for len(queue) > 0 {
		x, y, step := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if (x == 0 || x == rows-1 || y == 0 || y == cols-1) && !(x == entrance[0] && y == entrance[1]) {
			return step
		}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] && maze[nx][ny] == '.' {
				visited[nx][ny] = true
				queue = append(queue, []int{nx, ny, step + 1})
			}
		}
	}
	return -1
}

func main() {
	maze := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}
	entrance := []int{1, 2}
	fmt.Println(nearestExit(maze, entrance))
}
