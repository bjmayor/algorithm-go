# 🔋 LeetCode 2141：如何让 N 台电脑续航最久？——"二分答案"套路一文讲透

> 本文通过 LeetCode 2141 这道经典题目，讲解"最大-最小值（二分答案）"这一高频算法套路，并附带完整 Go 实现与单元测试。

---

## 题目描述

**2141. 同时运行 N 台电脑的最长时间**（[LeetCode 链接](https://leetcode.cn/problems/maximum-running-time-of-n-computers/)）

给你一个整数 `n`，表示有 `n` 台电脑，以及一个下标从 0 开始的整数数组 `batteries`，其中 `batteries[i]` 表示第 `i` 块电池可以提供的运行分钟数。

你可以在任意时刻将电池插入任意一台电脑。每次一块电池**只能供一台电脑使用**，但你可以**随时更换电池**（比如把电池从一台电脑上拔下来换到另一台）。当一块电池电量耗尽时它就不能再使用。

请返回在**最优电池分配方案**下，使得所有 `n` 台电脑**同时运行**的**最长分钟数**。

### 示例

| 输入 | 输出 | 说明 |
|------|------|------|
| n=2, batteries=[3,3,3] | 4 | 三块电池总电量 9，平均 4.5；通过合理更换可让两台电脑各跑 4 分钟 |
| n=2, batteries=[1,1,1,1] | 2 | 总电量 4，平均 2 |
| n=1, batteries=[10] | 10 | 单台电脑可串联使用所有电池 |

### 约束

- `1 <= n <= batteries.length <= 10^5`
- `1 <= batteries[i] <= 10^9`

---

## 引言

这道题的核心考查点是：如何在多块电池之间分配有限电量，使 n 台电脑在同一时刻能够运行尽可能久。实现上常见且高效的套路是**"在答案空间二分（最大-最小值/二分答案）"**，本文把这个套路讲清楚并给出易于复现的 Go 参考实现。

核心思路（为什么是“最大-最小值”）

目标是求最大的 t，使得所有 n 台电脑能同时运行 t 小时。把问题转化为判定问题：给定 t，判断能否分配电池使所有电脑同时运行 t 小时。构造判定函数：

`f(t) = sum_{i} min(batteries[i], t) >= n * t`

直观上，每块电池在满足单台电脑 t 小时时最多贡献 t 小时（因为它不能在同一时刻分给多台电脑），因此每块电池对目标 t 的贡献为 min(battery_i, t)。若所有电池的贡献之和不少于 n * t，就存在一种分配方式。

单调性保证二分可行：若 f(t) 为真，那么任意 t' <= t 也一定为真；若 f(t) 为假，那么任意 t' >= t 也为假。由此可以在区间上做二分搜索找最大 t。

上界选取：`high = sum(batteries) / n`（平均值）是一个安全的上界；下界 `low = 0`。

判定函数注意点

- 使用 64 位整型（int64）累加，避免总和溢出（最大 sum 可达 1e14）。
- min 操作越界或性能问题一般不用担心，但写法要简洁。
- 极端情况：n = 1 时，一台电脑可以串联使用所有电池，因此答案为 sum(batteries)。

伪代码

1. sum = sum(batteries)，low = 0，high = sum / n
2. while low < high:
   - mid = (low + high + 1) // 2    // 向上取整，寻找最大可行值
   - if check(mid): low = mid
   - else: high = mid - 1
3. 返回 low

check(t):
    total = 0
    for each b in batteries:
        total += min(b, t)
        if total >= `n * t`: return true
    return total >= `n * t`

复杂度

- 时间复杂度：每次判定 O(m)，二分的次数 O(log(sum/n))，总体 O(m * log S)，m 为电池数量，S 为 sum/n。
- 空间复杂度：O(1) 额外。

Go 参考实现（核心部分）

```go
func maxRunningTime(n int, batteries []int) int64 {
	var sum int64
	for _, b := range batteries {
		sum += int64(b)
	}
	low, high := int64(0), sum/int64(n)

	check := func(t int64) bool {
		var total int64
		for _, b := range batteries {
			if int64(b) >= t {
				total += t
			} else {
				total += int64(b)
			}
			if total >= t*int64(n) {
				return true
			}
		}
		return total >= t*int64(n)
	}

	for low < high {
		mid := (low + high + 1) / 2
		if check(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}
```

示例与边界测试

- 示例 1：n=2, batteries=[3,3,3] -> 输出 4
- 示例 2：n=2, batteries=[1,1,1,1] -> 输出 2
- 极端测试：n=1, batteries=[10] -> 输出 10（等于 sum）

可以把上面函数放到 `main.go` 中，并用若干测试用例进行验证。

参考与扩展阅读

- LeetCode 题目：2141 Maximum Running Time of N Computers
- 最大-最小值/二分答案套路详解（CSDN）：https://blog.csdn.net/goodparty/article/details/154545016

总结

这类题目非常适合用“二分答案”来解决：我们不是在数组中找特定元素，而是在“答案空间”二分；关键在于写出可判断某个答案是否可行的判定函数，并证明该判定是单调的。掌握这种思维后，面对很多“求最大的可行值”或“最小满足条件的值”问题会变得自然且高效。
