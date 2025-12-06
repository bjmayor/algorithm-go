# LeetCode 3578：统计极差最大为 K 的分割方式数 - 深入浅出指南

## 难度评析：为什么这是道"隐藏的困难题"？

LeetCode 标记这道题为**中等题**，但实际难度远超预期。原因如下：

| 维度 | 说明 |
|------|------|
| **问题理解** | 容易理解，但找到**有效的分割范围很难** |
| **思路难度** | 需要认识到"满足条件的前置位置形成连续区间"这个非平凡的性质 |
| **实现难度** | 虽然代码不复杂，但要想到"从右往左遍历"这个技巧很不直观 |
| **容易犯的错误** | 直接暴力枚举区间并现算极差是 O(n³)；即便加上“极差超限即 break”的朴素 DP 也有最坏 O(n²)，可能超时 |

**比较同类题**：

- 常见中等题通常考单一知识点（如前缀和、二分搜索）
- 这道题需要结合 **DP + 区间单调性 + 贪心观察**，属于"算法综合"题
- 难度更接近 **LeetCode Hard** 的下限


## 问题分析

这道题要求我们将数组分割成若干个连续子段，每个子段内**最大值与最小值的差不超过 k**，求总的分割方式数。

### 为什么容易超时？

直观的想法可能是：

```text
对于每个位置 i，枚举所有可能的前一个分割点 j，
检查 nums[j..i] 是否满足条件。
```

这样做的问题是：

- 对每个位置 i，我们需要检查 O(n) 个前置位置
- 对每个前置位置 j，检查子段 [j, i] 的最大值和最小值需要 O(n) 时间
- **总时间复杂度 O(n³)，肯定超时**


## 解题思路：DP + 滑动窗口（单调队列）

关键观察：**对于位置 i，满足条件的前置位置 j 不是分散的，而是一个连续的区间**。用滑动窗口维护这个区间的极差即可把复杂度压到 O(n)。

### 为什么？

假设位置 i 是当前分割点，我们要找所有有效的前一个分割点。

当我们从右往左扫描时：

- 如果 `nums[j..i]` 的极差 > k，那么 `nums[j-1..i]` 的极差也 > k（只会更大）
- 因此，满足条件的 j 会形成一个**连续的右侧区间**

例如，如果 j = [2, 3, 4] 都满足，那么 j = 1 就不满足。


### 算法框架（O(n)）

```text
1. 定义 dp[i]：前 i 个元素的分割方式数；prefix[i] 为 dp[0..i] 的前缀和。
2. 维护滑动窗口 [l, r]，用两个单调队列维护窗口内最小值与最大值。
3. 当窗口极差 > k 时，右移 l 并同步弹出队头下标。
4. 对于右端点 r：合法起点是区间 [l, r]，对应 dp 累加区间 dp[l..r]。
    dp[r+1] = prefix[r] - prefix[l-1]（注意取模）。
5. prefix[r+1] = prefix[r] + dp[r+1]，继续扩展 r。
```


## 代码实现

### 朴素 DP（O(n²)）

```go
// 逐步扩展左端点，遇到极差超出立刻 break，适合一般数据但最坏 O(n²)
func countPartitionsV1(nums []int, k int) int {
    const MOD = int(1e9 + 7)
    n := len(nums)
    dp := make([]int, n+1)
    dp[0] = 1
    for i := 0; i < n; i++ {
        minVal, maxVal := nums[i], nums[i]
        for j := i; j >= 0; j-- {
            if nums[j] < minVal {
                minVal = nums[j]
            }
            if nums[j] > maxVal {
                maxVal = nums[j]
            }
            if maxVal-minVal > k {
                break
            }
            dp[i+1] = (dp[i+1] + dp[j]) % MOD
        }
    }
    return dp[n]
}
```

### 优化版：滑动窗口 + 单调队列 + 前缀和（O(n)）

```go
package main

import "fmt"

const MOD = int(1e9 + 7)

func countPartitions(nums []int, k int) int {
    n := len(nums)
    dp := make([]int, n+1)
    prefix := make([]int, n+1)
    dp[0] = 1
    prefix[0] = 1

    minQ, maxQ := make([]int, 0, n), make([]int, 0, n)
    l := 0

    for r, v := range nums {
        for len(maxQ) > 0 && nums[maxQ[len(maxQ)-1]] <= v {
            maxQ = maxQ[:len(maxQ)-1]
        }
        maxQ = append(maxQ, r)

        for len(minQ) > 0 && nums[minQ[len(minQ)-1]] >= v {
            minQ = minQ[:len(minQ)-1]
        }
        minQ = append(minQ, r)

        for len(maxQ) > 0 && len(minQ) > 0 && nums[maxQ[0]]-nums[minQ[0]] > k {
            if maxQ[0] == l {
                maxQ = maxQ[1:]
            }
            if minQ[0] == l {
                minQ = minQ[1:]
            }
            l++
        }

        leftPrefix := 0
        if l > 0 {
            leftPrefix = prefix[l-1]
        }
        dp[r+1] = (prefix[r] - leftPrefix + MOD) % MOD
        prefix[r+1] = (prefix[r] + dp[r+1]) % MOD
    }

    return dp[n]
}

func main() {
    fmt.Println(countPartitions([]int{9, 4, 1, 3, 7}, 4)) // 6
    fmt.Println(countPartitions([]int{3, 3, 4}, 0))       // 2
}
```

## 时间复杂度分析

使用滑动窗口 + 单调队列，每个元素只进出队一次：

- 维护窗口极差的操作均摊 O(1)
- 计算 dp[r+1] 用前缀和 O(1)
- 总体时间复杂度 O(n)，空间 O(n)


## 实现细节与坑点

- 左指针右移时要同步弹出队头下标，否则窗口极差判断会错误。
- 计算 `dp[r+1]` 用前缀和差值时记得取模，并处理 `l=0` 的边界。
- 单调队列里存下标，不存值，便于与滑动窗口同步。


## 核心要点总结

| 要素 | 说明 |
|------|------|
| **关键观察** | 满足条件的前置位置形成连续区间 |
| **DP定义** | `dp[i]` = 以位置 i-1 结尾的分割方式数 |
| **窗口策略** | 单调队列维护窗口极差，超出 k 时右移左端点 |
| **时间复杂度** | O(n) |
| **空间复杂度** | O(n) |

## 测试验证（O(n) 实现）

用滑动窗口版计算 `nums = [9,4,1,3,7], k = 4`：

```text
初始化：dp[0]=1, prefix[0]=1, l=0, minQ=[], maxQ=[]

r=0, v=9: 窗口[0,0] 极差=0≤4
    dp[1] = prefix[0] - prefix[-1]=1
    prefix[1]=2

r=1, v=4: 窗口[0,1] 极差=5>4 → l 移到 1
    有效窗口[1,1]
    dp[2] = prefix[1] - prefix[0] = 2-1 = 1
    prefix[2]=3

r=2, v=1: 窗口[1,2] 极差=3≤4
    dp[3] = prefix[2] - prefix[0] = 3-1 = 2
    prefix[3]=5

r=3, v=3: 窗口[1,3] 极差=3≤4
    dp[4] = prefix[3] - prefix[0] = 5-1 = 4
    prefix[4]=9

r=4, v=7: 窗口[1,4] 极差=6>4，移 l→2，仍超，移 l→3
    有效窗口[3,4]
    dp[5] = prefix[4] - prefix[2] = 9-3 = 6
    prefix[5]=15

最终答案 dp[5]=6，与题目输出一致。
```

## 总结

这道题的核心难点不在于算法复杂性，而在于**找到合适的起点**。与其枚举所有前置位置然后验证，不如利用**极差单调性**从右往左扫描，通过及时 break 来避免不必要的计算。

这样的思路在许多数组分割问题中都适用。记住：**观察约束条件的单调性，往往能带来指数级的性能提升！**

