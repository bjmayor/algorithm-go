# LeetCode 3652: 按策略买卖股票的最佳时机

## 题目理解

给定价格数组 `prices` 和策略数组 `strategy`，策略可以是：
- `-1`: 买入
- `0`: 持有  
- `1`: 卖出

利润 = Σ(strategy[i] × prices[i])

我们可以进行**最多一次修改**：选择连续 k 个元素，前 k/2 个改为 0，后 k/2 个改为 1。

求最大可能利润。

## 解题思路

### 方法一：暴力枚举（朴素思路）

最直接的想法是枚举所有可能的修改位置：
- 不修改的情况
- 从索引 0, 1, 2, ..., n-k 开始修改

每次计算修改后的总利润，取最大值。

**时间复杂度**: O(n²) - 枚举 O(n) 个位置，每次重新计算总利润 O(n)

### 方法二：前缀和优化（推荐）

观察到暴力方法有大量重复计算。我们可以用**前缀和**优化。

#### 核心思想

设原始利润为 `base = Σ(strategy[i] × prices[i])`

当我们修改从位置 i 开始的 k 个元素时：
```
修改前: [i, i+1, ..., i+k/2-1, i+k/2, ..., i+k-1]
修改后: [     全部变成0      ][    全部变成1     ]
```

设：
- `p1 = i`, `p2 = i + k/2 - 1`（前半部分，变成 0）
- `p3 = i + k/2`, `p4 = i + k - 1`（后半部分，变成 1）

则：
- 原利润在 [p1, p2] 区间: `base1 = Σ(strategy[j] × prices[j])`，修改后变为 0
- 原利润在 [p3, p4] 区间: `base2 = Σ(strategy[j] × prices[j])`，修改后变为 `Σprices[j]`

**关键公式**：
```
新利润 = base - base1 - base2 + Σprices[p3~p4]
```

即：
```
profit_diff = Σprices[p3~p4] - base1 - base2
```

只有当 `profit_diff > 0` 时，修改才能提升利润。

#### 前缀和预处理

使用两个前缀和数组：
1. **prefixSums[i]**: prices 的前缀和，用于快速计算价格区间和
2. **baseSums[i]**: strategy[j] × prices[j] 的前缀和，用于快速计算原利润

这样每次查询区间和的时间从 O(k) 降到 O(1)。

**时间复杂度**: O(n) - 预处理 O(n)，枚举 O(n) 个位置，每次 O(1) 计算

**空间复杂度**: O(n)

## 代码实现

```go
func maxProfit(prices []int, strategy []int, k int) int64 {
    n := len(prices)
    
    // 前缀和预处理
    prefixSums := make([]int64, n+1)  // prices 的前缀和
    baseSums := make([]int64, n+1)    // strategy[i]*prices[i] 的前缀和
    var base int64
    
    for i := 0; i < n; i++ {
        prefixSums[i+1] = prefixSums[i] + int64(prices[i])
        base += int64(strategy[i]) * int64(prices[i])
        baseSums[i+1] = baseSums[i] + int64(strategy[i])*int64(prices[i])
    }
    
    // 边界：k > n 时无法修改
    if k > n {
        return base
    }
    
    maxProfit := base
    
    // 枚举所有修改起点
    for i := 0; i <= n-k; i++ {
        p1 := i
        p2 := i + k/2 - 1
        p3 := i + k/2
        p4 := i + k - 1
        
        // 原利润中被修改部分的贡献
        base1 := baseSums[p2+1] - baseSums[p1]
        base2 := baseSums[p4+1] - baseSums[p3]
        
        // 修改后后半部分的贡献（全为1）
        priceSum := prefixSums[p4+1] - prefixSums[p3]
        
        // 新利润 = 原利润 - 被移除部分 + 新增部分
        profit := base - base1 - base2 + priceSum
        
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    
    return maxProfit
}
```

## 示例演示

以 `prices = [4,2,8], strategy = [-1,0,1], k = 2` 为例：

**预处理**：
- base = (-1)×4 + 0×2 + 1×8 = 4
- prefixSums = [0, 4, 6, 14]
- baseSums = [0, -4, -4, 4]

**枚举修改位置**：
- i=0: 修改 [0,1] → [0,1,1]
  - base1 = baseSums[1] - baseSums[0] = -4
  - base2 = 0 (p3=p4+1)
  - priceSum = prefixSums[2] - prefixSums[1] = 2
  - profit = 4 - (-4) - 0 + 2 = **10** ✓

- i=1: 修改 [1,2] → [-1,0,1] (无变化)
  - profit = 4

**最大利润**: 10

## 复杂度分析

- **时间复杂度**: O(n)，其中 n 是数组长度
- **空间复杂度**: O(n)，用于存储前缀和数组

## 总结

本题的关键是识别暴力枚举中的重复计算，通过前缀和实现 O(1) 的区间和查询，将时间复杂度从 O(n²) 优化到 O(n)。这是一个典型的「暴力→优化」的思维过程。




