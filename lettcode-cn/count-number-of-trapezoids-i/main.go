package main

const MOD = 1_000_000_007

// countTrapezoids 统计可以从 points 中选择四个不同点组成的水平梯形数量
// 水平梯形：具有至少一对平行于 x 轴的边的凸四边形
func countTrapezoids(points [][]int) int {
	// 按 y 坐标分组，统计每个 y 值出现的次数
	yCount := make(map[int]int)
	for _, p := range points {
		yCount[p[1]]++
	}

	// 计算每条水平线上的边对数 C(cnt, 2)
	// 然后用累加技巧计算所有不同行的组合
	var sum, ans int64 = 0, 0

	for _, cnt := range yCount {
		if cnt < 2 {
			continue
		}
		// C(cnt, 2) = cnt * (cnt - 1) / 2
		pairs := int64(cnt) * int64(cnt-1) / 2 % MOD
		ans = (ans + pairs*sum) % MOD
		sum = (sum + pairs) % MOD
	}

	return int(ans)
}

func main() {
	// 示例测试
	points1 := [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}}
	println("示例1:", countTrapezoids(points1)) // 期望输出: 3

	points2 := [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}}
	println("示例2:", countTrapezoids(points2)) // 期望输出: 1
}
