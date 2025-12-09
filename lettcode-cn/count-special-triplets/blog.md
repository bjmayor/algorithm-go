# LeetCode 3583: 统计特殊三元组 - 从 O(n³) 到 O(n) 的顿悟之旅

## 题目回顾

给定整数数组 `nums`，找出满足以下条件的三元组 `(i, j, k)` 的数量：
- `0 <= i < j < k < n`
- `nums[i] == nums[j] * 2`
- `nums[k] == nums[j] * 2`

## 第一反应：暴力三重循环

看到这道题，你的第一反应可能是：

```go
func specialTriplets(nums []int) int {
    count := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                if nums[i] == nums[j]*2 && nums[k] == nums[j]*2 {
                    count++
                }
            }
        }
    }
    return count
}
```

**时间复杂度：O(n³)**

这显然不够优雅。当 `n = 10⁵` 时，这种做法会超时。

## 关键洞察：固定中间，向两边看

💡 **顿悟时刻 #1**：三元组问题的本质

仔细观察条件：
- `nums[i] == nums[j] * 2`（i 在 j 左边）
- `nums[k] == nums[j] * 2`（k 在 j 右边）

**关键发现**：两个条件都依赖于 `nums[j]`！

如果我们**固定 j 作为中间点**，问题就变成了：
- 在 j 左边有多少个值等于 `nums[j] * 2`？
- 在 j 右边有多少个值等于 `nums[j] * 2`？

答案就是：**左边数量 × 右边数量**（乘法原理）

## 💡 顿悟时刻 #2：用空间换时间

如何快速知道左边和右边有多少个目标值？

**答案：频率表（哈希表）！**

维护两个频率表：
- `freqPrev`：记录 j 左边各个值出现的次数
- `freqNext`：记录 j 右边各个值出现的次数

查询 `freqPrev[nums[j] * 2]` 和 `freqNext[nums[j] * 2]` 都是 **O(1)** 操作！

## 💡 顿悟时刻 #3：动态维护是关键

这里是最精妙的地方！

你可能会想：每个 j 都要重新计算左右的频率表？那不还是 O(n²) 吗？

**不！我们可以动态维护！**

```
初始状态：
freqPrev = {}（空，因为 j=0 左边没有元素）
freqNext = 整个数组的频率统计

当 j 向右移动一位时：
1. nums[j] 从"右边"变成了"当前"
2. 下一轮，nums[j] 就在"左边"了

所以只需要：
freqNext[nums[j]]--  // 从右边移除
freqPrev[nums[j]]++  // 加入左边
```

**每次更新只需 O(1)！**

## 完整实现

```go
func specialTriplets(nums []int) int {
    const MOD = 1_000_000_007
    tripletCount := 0
    n := len(nums)
    
    // 初始化频率表
    freqPrev := make(map[int]int)
    freqNext := make(map[int]int)
    
    // freqPrev 初始只包含 nums[0]
    freqPrev[nums[0]]++
    
    // freqNext 包含 nums[1:] 的所有元素
    for i := 1; i < n; i++ {
        freqNext[nums[i]]++
    }
    
    // 遍历中间点 j（从 1 到 n-2）
    for j := 1; j < n-1; j++ {
        // 先从 freqNext 中移除 nums[j]（j 变成当前点了）
        freqNext[nums[j]]--
        if freqNext[nums[j]] == 0 {
            delete(freqNext, nums[j])
        }
        
        // 计算贡献（注意取模）
        target := nums[j] * 2
        if freqPrev[target] > 0 && freqNext[target] > 0 {
            contribution := (freqPrev[target] * freqNext[target]) % MOD
            tripletCount = (tripletCount + contribution) % MOD
        }
        
        // 将 nums[j] 加入 freqPrev（为下一轮做准备）
        freqPrev[nums[j]]++
    }
    
    return tripletCount
}
```

## 复杂度分析

**时间复杂度：O(n)**

- 初始化 freqNext：O(n)
- 主循环：n 次迭代，每次 O(1) 操作
- 总计：O(n)

**空间复杂度：O(n)**

- 两个哈希表，最坏情况存储所有不同的值

## ⚠️ 注意事项：取模运算

题目要求结果对 `10⁹ + 7` 取模，因为：

- `n` 最大可达 `10⁵`
- 最坏情况下，所有元素都相同，三元组数量可能达到 $\binom{n}{3} \approx \frac{n^3}{6}$
- 当 `n = 10⁵` 时，结果可能超过 `int` 范围

**关键点**：每次累加贡献时都要取模

```go
contribution := (freqPrev[target] * freqNext[target]) % MOD
tripletCount = (tripletCount + contribution) % MOD
```

这样可以防止中间计算溢出！

## 核心思想总结

1. **固定中间**：将三元组问题转化为"固定中间点，向两边统计"
2. **频率表**：用哈希表 O(1) 查询左右区间的目标值数量
3. **动态维护**：滑动窗口思想，边遍历边更新频率表，避免重复计算

从 O(n³) 到 O(n)，关键在于**找到问题的对称性和可复用性**！

## 举例验证

对于 `nums = [8, 4, 2, 8, 4]`：

```
j=1, nums[j]=4, target=8:
  freqPrev: {8: 1}
  freqNext: {2: 1, 8: 1, 4: 1}
  贡献: 1 × 1 = 1  ✓ (i=0, j=1, k=3)

j=2, nums[j]=2, target=4:
  freqPrev: {8: 1, 4: 1}
  freqNext: {8: 1, 4: 1}
  贡献: 1 × 1 = 1  ✓ (i=1, j=2, k=4)

总计: 2 ✓
```

完美！🎯
