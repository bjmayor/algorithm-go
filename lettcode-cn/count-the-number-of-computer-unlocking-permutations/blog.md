
# 统计计算机解锁顺序排列数 — 题解与思路

这道题表面看起来像图遍历或回溯问题，但通过观察约束与解锁规则，可以得到非常简单直接的结论。下面以清晰的思路、证明以及代码实现来解释本题。

问题回顾：

- 给定长度为 n 的数组 `complexity`，编号为 `0..n-1` 的计算机每台有一个密码复杂度 `complexity[i]`。
- 编号为 `0` 的计算机一开始已解锁（root）。其他计算机 `i` 需要通过某个 `j < i 且 complexity[j] < complexity[i]` 的已解锁机器解锁。
- 问：有多少个排列可以表示从 `0` 开始解锁所有机器的顺序？答案模 10^9+7。

关键观察（结论）：

- 若 `complexity[0]` 是数组中的**唯一最小值**，则任何 i>0 都能被 0 解锁，因此其余 n-1 台机器可以按任意顺序解锁，答案是 `(n-1)!`（模 10^9+7）。
- 否则（存在 `i>0` 使 `complexity[i] <= complexity[0]`），则一定存在某个最早的索引 t，满足 `complexity[t] <= complexity[0]`，此时 t 无法被其之前的任何机器解锁（没有 j < t 且 complexity[j] < complexity[t]），因此答案为 0。

证明（必要性与充分性，简要）：

- 充分性：若 `complexity[0]` 严格小于其余所有元素，那么 0 可以直接解锁任意 i>0。由于没有其它约束，剩余 n-1 个机器的排列皆可行，所以有 `(n-1)!` 种。
- 必要性：若存在 `t>0` 且 `complexity[t] <= complexity[0]`（取最小的 t），则对任何 j < t，`complexity[j] >= complexity[t]`（否则 t 不是最早违反条件的位置）。因此没有 j < t 且 `complexity[j] < complexity[t]`，t 无法解锁，违背题意。所以无解、返回 0。

复杂度分析：

- 时间复杂度：O(n)，只需遍历一次数组确认 `complexity[0]` 是否为唯一最小值
- 空间复杂度：O(1)

示例：

- `complexity = [1,2,3]` → `2`（(3-1)! = 2）
- `complexity = [3,3,3,4,4,4]` → `0`
- `complexity = [1,2,2]` → `2`（`0` 可解锁所有，余下两台任意排列）
- `complexity = [1,1,2]` → `0`（第二个 `1` 无法被更小的索引解锁）

Go 实现：

```go
package main

import "fmt"

const MOD = 1000000007

// countPermutations 返回能解锁所有计算机的排列数（模 MOD）。
// 若 complexity[0] 不是严格小于其余元素（即不是唯一的最小值），返回 0。
func countPermutations(complexity []int) int {
	n := len(complexity)
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}
	ans := 1
	for i := 2; i <= n-1; i++ {
		ans = (ans * i) % MOD
	}
	return ans
}

func main() {
	examples := [][]int{{1,2,3}, {3,3,3,4,4,4}}
	for _, ex := range examples {
		fmt.Printf("input: %v, output: %d\n", ex, countPermutations(ex))
	}
}
```

常见误区：
- 误以为题目需要复杂的回溯或图遍历——题目限制 `j < i` 且 `complexity[j] < complexity[i]` 使得全局解构变得简单可判定。
- 误以为数组中不能有重复值才有解 —— 实际上只要 `complexity[0]` 严格最小就行，数组中其他位置可以重复。

结语：
这道题是通过观察约束找到决定性条件的典型示例：一旦发现 `complexity[0]` 是否为唯一最小值就能直接判定结果并给出公式化解法，避免了不必要的复杂性。


