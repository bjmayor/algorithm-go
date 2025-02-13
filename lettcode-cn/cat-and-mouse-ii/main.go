package main // 定义主包

import "fmt" // 导入格式化输出包

/**
 * @title 1728. 猫和老鼠 II
 * @link https://leetcode.cn/problems/cat-and-mouse-ii/
 *
 * @description
 * 一只猫和一只老鼠在一个网格中玩游戏。网格中有墙壁、空地、猫的起始位置、老鼠的起始位置和食物。
 *
 * @rules 游戏规则
 * 1. 老鼠和猫轮流移动,老鼠先手
 * 2. 每次移动可以上下左右移动1-jumpLength个格子,但不能穿过墙壁
 * 3. 游戏结束条件:
 *    - 如果猫和老鼠处于同一个单元格,猫获胜
 *    - 如果老鼠到达食物的位置,老鼠获胜
 *    - 如果游戏达到1000轮仍未结束,老鼠获胜
 * 4. 假设双方都采用最优策略
 *
 * @solution 解题思路
 * 1. 使用记忆化搜索
 * 2. 状态定义: dp[mouseX][mouseY][catX][catY][turns] 表示在特定状态下的游戏结果
 * 3. 状态转移:
 *    - 老鼠回合:尝试所有可能移动,如果存在一种移动可以获胜,则选择该移动
 *    - 猫回合:尝试所有可能移动,如果存在一种移动可以获胜,则选择该移动
 * 4. 返回值:
 *    MOUSE_WIN(1): 老鼠获胜
 *    CAT_WIN(2): 猫获胜
 *    DRAW(0): 平局
 *
 * @complexity
 * - 时间复杂度: O(R * C * R * C * MAX_MOVES)，其中 R 和 C 分别是网格的行数和列数
 * - 空间复杂度: O(R * C * R * C * MAX_MOVES)，用于存储记忆化数组
 */

// 定义游戏结果常量
const (
	MOUSE_WIN = 1    // 老鼠获胜
	CAT_WIN   = 2    // 猫获胜
	DRAW      = 0    // 平局
	MAX_MOVES = 1000 // 最大移动次数
)

// 方向数组,用于上下左右移动
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上、下、左、右四个方向

// 主要解题函数
func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows, cols := len(grid), len(grid[0]) // 获取网格的行数和列数
	var mouseStart, catStart, food []int  // 定义起始位置和食物位置

	// 找到起始位置和食物位置
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'M' { // 找到老鼠起始位置
				mouseStart = []int{i, j}
			} else if grid[i][j] == 'C' { // 找到猫起始位置
				catStart = []int{i, j}
			} else if grid[i][j] == 'F' { // 找到食物位置
				food = []int{i, j}
			}
		}
	}

	// 创建记忆化数组
	// dp[mouseX][mouseY][catX][catY][turns] 表示:
	// - mouseX,mouseY: 老鼠的位置坐标
	// - catX,catY: 猫的位置坐标  
	// - turns: 当前回合数
	// 返回值表示在该状态下的游戏结果
	dp := make([][][][][]int, rows) // 创建5维数组
	for i := range dp {
		dp[i] = make([][][][]int, cols)
		for j := range dp[i] {
			dp[i][j] = make([][][]int, rows)
			for k := range dp[i][j] {
				dp[i][j][k] = make([][]int, cols)
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = make([]int, MAX_MOVES)
					for m := range dp[i][j][k][l] {
						dp[i][j][k][l][m] = -1 // 初始化为-1表示未访问
					}
				}
			}
		}
	}

	// DFS函数声明
	var dfs func(mouseX, mouseY, catX, catY, turns int) int
	dfs = func(mouseX, mouseY, catX, catY, turns int) int {
		if turns >= MAX_MOVES { // 超过最大移动次数,老鼠获胜
			return MOUSE_WIN
		}
		if mouseX == catX && mouseY == catY { // 猫和老鼠在同一位置,猫获胜
			return CAT_WIN
		}
		if mouseX == food[0] && mouseY == food[1] { // 老鼠到达食物位置,老鼠获胜
			return MOUSE_WIN
		}
		if catX == food[0] && catY == food[1] { // 猫到达食物位置,猫获胜
			return CAT_WIN
		}

		if dp[mouseX][mouseY][catX][catY][turns] != -1 { // 如果该状态已经计算过,直接返回
			return dp[mouseX][mouseY][catX][catY][turns]
		}

		var result int
		if turns%2 == 0 { // 老鼠回合
			result = CAT_WIN           // 初始假设猫获胜
			for _, dir := range dirs { // 遍历四个方向
				for step := 0; step <= mouseJump; step++ { // 在跳跃范围内尝试每一步
					newMouseX, newMouseY := mouseX+dir[0]*step, mouseY+dir[1]*step                                                     // 计算新位置
					if newMouseX < 0 || newMouseX >= rows || newMouseY < 0 || newMouseY >= cols || grid[newMouseX][newMouseY] == '#' { // 检查是否越界或撞墙
						break
					}
					nextResult := dfs(newMouseX, newMouseY, catX, catY, turns+1) // 递归计算下一步
					if nextResult == MOUSE_WIN {                                 // 如果找到一种获胜方式
						result = MOUSE_WIN
						break
					}
				}
				if result == MOUSE_WIN { // 如果已经找到获胜方式,退出循环
					break
				}
			}
		} else { // 猫回合
			result = MOUSE_WIN         // 初始假设老鼠获胜
			for _, dir := range dirs { // 遍历四个方向
				for step := 0; step <= catJump; step++ { // 在跳跃范围内尝试每一步
					newCatX, newCatY := catX+dir[0]*step, catY+dir[1]*step                                                 // 计算新位置
					if newCatX < 0 || newCatX >= rows || newCatY < 0 || newCatY >= cols || grid[newCatX][newCatY] == '#' { // 检查是否越界或撞墙
						break
					}
					nextResult := dfs(mouseX, mouseY, newCatX, newCatY, turns+1) // 递归计算下一步
					if nextResult == CAT_WIN {                                   // 如果找到一种获胜方式
						result = CAT_WIN
						break
					}
				}
				if result == CAT_WIN { // 如果已经找到获胜方式,退出循环
					break
				}
			}
		}

		dp[mouseX][mouseY][catX][catY][turns] = result // 记录当前状态的结果
		return result                                  // 返回结果
	}

	// 从初始状态开始搜索
	result := dfs(mouseStart[0], mouseStart[1], catStart[0], catStart[1], 0)
	return result == MOUSE_WIN // 返回是否老鼠获胜
}

func main() {
	// 测试用例1: 老鼠可以获胜
	fmt.Println(canMouseWin([]string{"####F", "#C...", "M...."}, 1, 2)) // true
	// 测试用例2: 老鼠可以获胜
	fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 4)) //true
	// 测试用例3: 老鼠无法获胜
	fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 3)) //false
	// 测试用例4: 老鼠无法获胜
	fmt.Println(canMouseWin([]string{"C...#", "...#F", "....#", "M...."}, 2, 5)) // false
	// 测试用例5: 老鼠可以获胜
	fmt.Println(canMouseWin([]string{".M...", "..#..", "#..#.", "C#.#.", "...#F"}, 3, 1)) // true
}
