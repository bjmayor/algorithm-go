package main

import "fmt"

/*
* uniquePathsWithObstacles 计算带障碍物的不同路径数
* 算法思想:
* 1. 使用动态规划,dp[j]表示到达位置(i,j)的不同路径数.dp[j] 表示到达当前行 i 的第 j 列的不同路径数。
* 由于我们只需要当前行和前一行的信息，所以可以使用一维数组 dp 来存储。
* 2. 状态转移:
*    - 如果当前位置无障碍(obstacleGrid[i][j] == 0):
*      dp[j] += dp[j-1] (从左边来的路径数)
*    - 如果当前位置有障碍(obstacleGrid[i][j] == 1):
*      dp[j] = 0 (无法到达该位置)
* 3. 初始化:
*    - 如果起点无障碍,dp[0] = 1
*    - 否则dp[0] = 0
*
* 时间复杂度: O(m*n) - 遍历整个网格
* 空间复杂度: O(n) - 使用一维dp数组
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 获取网格的行数和列数
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 创建dp数组,长度为列数
	dp := make([]int, n)
	// 初始化起点
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}
	// 遍历网格
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果当前位置无障碍
			if obstacleGrid[i][j] == 0 {
				// 加上从左边来的路径数
				if j-1 >= 0 {
					dp[j] += dp[j-1]
				}
			} else {
				// 有障碍物,无法到达
				dp[j] = 0
			}
		}
	}
	// 返回终点的路径数
	return dp[n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})) //2
}
