# LeetCode 3562: 折扣价交易股票的最大利润 - 树形DP+背包问题详解

> 难度：困难  
> 标签：动态规划、树形DP、背包问题、DFS

## 一、题目描述

给定一个公司的员工层级关系（树形结构），每位员工都有购买自己股票的机会：
- `present[i]`：第 i 位员工今天购买股票的价格
- `future[i]`：第 i 位员工明天卖出股票的预期价格
- `hierarchy`：员工的上下级关系
- `budget`：总投资预算

**核心规则：**
1. 每只股票最多购买一次
2. 如果员工的直属上司购买了股票，该员工可以**半价**购买（floor(present[i] / 2)）
3. 不能用未来收益增加预算

**目标：** 在预算限制内，最大化总利润。

### 示例

**示例 1：**
```
输入：n = 3, present = [5,2,3], future = [8,5,6], 
     hierarchy = [[1,2],[2,3]], budget = 7
输出：12
解释：买员工1(花费5)→买员工2(半价1)→买员工3(半价1)
     利润=(8-5)+(5-1)+(6-1)=12
```

**示例 2：**
```
输入：n = 3, present = [4,6,8], future = [7,9,11], 
     hierarchy = [[1,2],[1,3]], budget = 10
输出：10
解释：买员工1(花费4)+买员工3(半价4)
     利润=(7-4)+(11-4)=10
```

## 二、问题分析

### 2.1 关键洞察

1. **树形结构**：员工层级关系构成树，需要树形DP
2. **预算约束**：有限预算下的最优化问题，需要背包DP
3. **折扣机制**：父节点购买影响子节点价格，状态需包含父节点信息
4. **多子节点**：需要在子节点间分配预算，使用分组背包

### 2.2 问题建模

这是一个 **树形DP + 分组背包** 的组合问题：
- **树形DP**：在树上做决策（每个节点买/不买）
- **分组背包**：多个子节点间分配有限预算

## 三、算法设计

### 3.1 状态定义

定义 `dfs(u)` 返回一个结果结构体：

```go
type result struct {
    dp0  []int  // dp0[b]: 父节点不购买时，预算b的最大收益
    dp1  []int  // dp1[b]: 父节点购买（当前节点可享半价）时，预算b的最大收益
    size int    // 子树最大可能开销（用于剪枝）
}
```

**状态含义：**
- `dp0[b]`：父节点**没买**，当前子树在预算 b 下的最大收益
- `dp1[b]`：父节点**买了**（当前节点可半价），预算 b 下的最大收益
- `size`：子树所有节点原价之和（用于背包剪枝）

**关键优化：** 一次 DFS 计算所有预算 [0, budget] 的结果，避免重复计算。

### 3.2 状态转移

对于每个节点 `u`，分两步计算：

**步骤1：分组背包合并所有子节点**

```go
// 初始化子节点收益数组
subProfit0 := make([]int, budget+1)  // 子节点无折扣
subProfit1 := make([]int, budget+1)  // 子节点有折扣（当前节点买了）

for _, child := range children[u] {
    childResult := dfs(child)
    
    // 分组背包：倒序遍历预算
    for i := budget; i >= 0; i-- {
        // 关键剪枝：只遍历到 min(childResult.size, i)
        for sub := 0; sub <= min(childResult.size, i); sub++ {
            subProfit0[i] = max(subProfit0[i], subProfit0[i-sub] + childResult.dp0[sub])
            subProfit1[i] = max(subProfit1[i], subProfit1[i-sub] + childResult.dp1[sub])
        }
    }
}
```

**步骤2：考虑当前节点的买/不买**

```go
cost := present[u]      // 原价
dCost := present[u] / 2 // 半价

for i := 0; i <= budget; i++ {
    // 父节点不购买的情况（dp0）
    dp0[i] = subProfit0[i]  // 不买当前节点
    if i >= cost {
        dp0[i] = max(dp0[i], subProfit1[i-cost] + future[u] - cost)  // 买当前节点(原价)
    }
    
    // 父节点购买的情况（dp1，当前节点可半价）
    dp1[i] = subProfit0[i]  // 不买当前节点
    if i >= dCost {
        dp1[i] = max(dp1[i], subProfit1[i-dCost] + future[u] - dCost)  // 买当前节点(半价)
    }
}
```

### 3.3 核心优化

**1. 批量计算所有预算**
- 原始做法：为每个预算值单独调用 DFS → O(n × budget³)
- 优化做法：一次返回所有预算的结果数组 → O(n × budget²)

**2. size 剪枝**
```go
for sub := 0; sub <= min(childResult.size, i); sub++ {
    // 子树最多花费 size，超过的预算遍历无意义
}
```

**3. 状态分离**
- 用 `dp0` 和 `dp1` 分别处理父节点买/不买的情况
- 避免用 boolean 参数递归，减少状态空间

## 四、代码实现

```go
type result struct {
    dp0  []int // 父节点不购买时，各预算下的最大收益
    dp1  []int // 父节点购买时（当前节点可享受半价），各预算下的最大收益
    size int   // 子树最大可能开销（用于剪枝）
}

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
    // 构建邻接表
    g := make([][]int, n)
    for _, edge := range hierarchy {
        parent, child := edge[0]-1, edge[1]-1
        g[parent] = append(g[parent], child)
    }

    var dfs func(u int) result
    dfs = func(u int) result {
        cost := present[u]      // 原价
        dCost := present[u] / 2 // 半价

        // dp0[b]: 父节点不购买，预算b的最大收益
        // dp1[b]: 父节点购买（可半价），预算b的最大收益
        dp0 := make([]int, budget+1)
        dp1 := make([]int, budget+1)

        // 子节点收益（分组背包）
        // subProfit0[b]: 子节点无折扣，预算b的最大收益
        // subProfit1[b]: 子节点有折扣（当前节点买了），预算b的最大收益
        subProfit0 := make([]int, budget+1)
        subProfit1 := make([]int, budget+1)

        uSize := cost // 当前子树最大开销（至少是当前节点的原价）

        // 分组背包：合并所有子节点
        for _, v := range g[u] {
            childResult := dfs(v)
            uSize += childResult.size // 累加子树开销

            // 倒序遍历预算，避免重复使用
            for i := budget; i >= 0; i-- {
                // 关键剪枝：只遍历到 min(childResult.size, i)
                for sub := 0; sub <= min(childResult.size, i); sub++ {
                    // 子节点无折扣
                    subProfit0[i] = max(subProfit0[i], subProfit0[i-sub]+childResult.dp0[sub])
                    // 子节点有折扣（当前节点购买）
                    subProfit1[i] = max(subProfit1[i], subProfit1[i-sub]+childResult.dp1[sub])
                }
            }
        }

        // 计算当前节点的 dp0 和 dp1
        for i := 0; i <= budget; i++ {
            // 父节点不购买的情况
            dp0[i] = subProfit0[i] // 不买当前节点
            if i >= cost {
                // 买当前节点（原价）+ 子节点享受折扣
                dp0[i] = max(dp0[i], subProfit1[i-cost]+future[u]-cost)
            }

            // 父节点购买的情况（当前节点可半价）
            dp1[i] = subProfit0[i] // 不买当前节点
            if i >= dCost {
                // 买当前节点（半价）+ 子节点享受折扣
                dp1[i] = max(dp1[i], subProfit1[i-dCost]+future[u]-dCost)
            }
        }

        return result{dp0, dp1, uSize}
    }

    return dfs(0).dp0[budget]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 五、示例推导

### 示例 1：n = 3, present = [5,2,3], future = [8,5,6], budget = 7

**树结构**：0 → 1 → 2（节点编号从0开始）

**叶子节点 2 (present=3, future=6)：**
```
cost = 3, dCost = 1
无子节点，subProfit0 和 subProfit1 全为 0

dp0[b]: 父节点不购买时的最大收益
  b < 3: 0 (买不起)
  b >= 3: 6-3 = 3 (买节点2，原价)

dp1[b]: 父节点购买时的最大收益（当前节点可半价）
  b < 1: 0 (买不起)
  b >= 1: 6-1 = 5 (买节点2，半价)

返回: {dp0=[0,0,0,3,3,3,3,3], dp1=[0,5,5,5,5,5,5,5], size=3}
```

**节点 1 (present=2, future=5)：**
```
cost = 2, dCost = 1
子节点 [2]: childResult = {dp0=[0,0,0,3,3,3,3,3], dp1=[0,5,5,5,5,5,5,5], size=3}

分组背包合并子节点：
  subProfit0[b] = 子节点无折扣时的收益 = [0,0,0,3,3,3,3,3]
  subProfit1[b] = 子节点有折扣时的收益 = [0,5,5,5,5,5,5,5]

计算 dp0[b]: 父节点不购买
  不买节点1: dp0[b] = subProfit0[b]
  买节点1(原价2): dp0[b] = max(subProfit0[b], subProfit1[b-2] + 5-2)
  
  dp0[2] = max(0, 5+3) = 8  ✓ (买节点1，子节点2半价收益5)
  dp0[3] = max(3, 5+3) = 8
  ...

计算 dp1[b]: 父节点购买（当前节点可半价）
  买节点1(半价1): dp1[b] = max(subProfit0[b], subProfit1[b-1] + 5-1)
  
  dp1[1] = max(0, 0+4) = 4 ❌
  dp1[1] = max(0, 5+4) = 9 ✓ (买节点1半价，子节点2也半价)

返回: {dp0=[0,0,8,8,8,8,8,8], dp1=[0,9,9,9,9,9,9,9], size=5}
```

**根节点 0 (present=5, future=8)：**
```
cost = 5, dCost = 2
子节点 [1]: childResult 的 dp0=[0,0,8,8,8,8,8,8], dp1=[0,9,9,9,9,9,9,9]

分组背包合并：
  subProfit0[b] = [0,0,8,8,8,8,8,8]  (子节点无折扣)
  subProfit1[b] = [0,9,9,9,9,9,9,9]  (子节点有折扣)

计算 dp0[7]: 父节点不购买
  不买节点0: dp0[7] = subProfit0[7] = 8
  买节点0(原价5): dp0[7] = max(8, subProfit1[7-5] + 8-5) = max(8, 9+3) = 12 ✓

最终答案: 12
```

**最优策略**：买节点0(花费5) → 买节点1(半价1) → 买节点2(半价1)  
**总花费**：5 + 1 + 1 = 7  
**总收益**：(8-5) + (5-1) + (6-1) = 12

### 示例 2：n = 3, present = [4,6,8], future = [7,9,11], budget = 10

**树结构**：0 → [1, 2]

**叶子节点：**
- 节点 1: {dp0中b≥6时为3, dp1中b≥3时为6, size=6}
- 节点 2: {dp0中b≥8时为3, dp1中b≥4时为7, size=8}

**根节点 0：**
```
分组背包合并两个子节点[1,2]：
  先加入节点1的收益，再加入节点2的收益
  
  subProfit0[10]: 不给子节点折扣
    = max(买节点1原价6得3, 买节点2原价8得3) = 3
  
  subProfit1[10]: 给子节点折扣（当前节点买了）
    = 考虑节点1半价3得6, 节点2半价4得7
    = 买节点2半价4得7 (预算6剩余可再买节点1半价3) = 7+6 = 13 ❌
    = 只买节点2半价4得7 ✓

计算 dp0[10]:
  不买节点0: 3
  买节点0(原价4): subProfit1[10-4] + 7-4 = subProfit1[6] + 3
    subProfit1[6] = 7 (只买节点2半价4)
    = 7 + 3 = 10 ✓

最终答案: 10
```

**最优策略**：买节点0(花费4) → 买节点2(半价4)  
**总花费**：4 + 4 = 8  
**总收益**：(7-4) + (11-4) = 10

## 六、复杂度分析

### 时间复杂度：O(n × budget²)
- n 个节点，每个节点调用一次 DFS
- 每个节点的背包DP：外层 O(budget)，内层最多 O(size) ≈ O(budget)
- 有 size 剪枝，实际复杂度更优

### 空间复杂度：O(budget)
- 每个节点需要 4 个 O(budget) 的数组：dp0, dp1, subProfit0, subProfit1
- 递归栈深度 O(n)
- 总空间 O(budget + n)

## 七、性能优化详解

### 7.1 朴素解法的问题

**朴素想法**：`dfs(u, budget, parentBought) -> int`

```go
// 为每个预算值单独计算
dfs := func(u, remainBudget int, parentBought bool) int {
    // 1. 不买当前节点
    result = knapsack(children, remainBudget, false)
    
    // 2. 买当前节点
    result = max(result, profit + knapsack(children, remainBudget-cost, true))
    
    return result
}
```

**问题**：
1. 每次调用只返回单个预算值的结果
2. knapsack 内部需要遍历所有子节点的所有预算组合
3. 大量重复计算

**时间复杂度**：O(n × budget³)

### 7.2 优化策略

**核心思路**：一次计算所有预算的结果

```go
// 一次返回所有预算[0..budget]的结果
dfs := func(u int) result {
    dp0 := make([]int, budget+1)  // 所有预算的结果
    dp1 := make([]int, budget+1)
    
    // 通过背包DP一次性计算所有预算
    // ...
    
    return result{dp0, dp1, size}
}
```

**优势**：
1. 避免重复计算同一节点的不同预算
2. 背包DP自然地处理所有预算
3. size 剪枝进一步减少搜索空间

### 7.3 性能对比

| 指标 | 朴素解法 | 优化解法 | 提升 |
|------|---------|---------|------|
| 时间复杂度 | O(n×budget³) | O(n×budget²) | budget 倍 |
| 空间复杂度 | O(n×budget×2) | O(budget) | n×2 倍 |
| 重复计算 | 大量 | 几乎无 | - |
| 剪枝效果 | 无 | size 剪枝 | 显著 |

**测试用例对比**（n=12, budget=93）：
- 朴素解法：约 12 × 93³ ≈ 9,600,000 次操作 → **超时**
- 优化解法：约 12 × 93² ≈ 104,000 次操作 → **通过**
- **性能提升约 92 倍**

## 八、知识回顾

### 8.1 0-1背包问题

#### 基本问题
有 n 个物品和容量为 W 的背包，每个物品重量 w[i]，价值 v[i]，每个物品最多选一次。

#### 状态转移
```
dp[i][w] = max(
    dp[i-1][w],              // 不选第i个物品
    dp[i-1][w-w[i]] + v[i]   // 选第i个物品
)
```

#### 空间优化（一维数组）
```go
dp := make([]int, W+1)
for i := 0; i < n; i++ {
    // 必须倒序遍历，避免重复使用
    for w := W; w >= weight[i]; w-- {
        dp[w] = max(dp[w], dp[w-weight[i]] + value[i])
    }
}
```

**倒序原因**：保证每个物品只用一次

#### 分组背包（本题使用）
物品分为若干组，每组最多选一个物品。

```go
for _, group := range groups {  // 遍历每组
    for w := W; w >= 0; w-- {   // 倒序遍历预算
        for _, item := range group {  // 遍历组内物品
            if w >= item.weight {
                dp[w] = max(dp[w], dp[w-item.weight] + item.value)
            }
        }
    }
}
```

**本题应用**：
- 每个子节点是一组
- 组内"物品"是给该子节点的不同预算分配
- `dfs(child, spend, parentBought)` 计算分配 spend 预算的价值

### 8.2 树形DP通用框架

#### 核心思想
在树结构上进行动态规划，通常从**叶子向根**或用**DFS递归**。

#### 通用模板
```go
func dfs(u int) State {
    // 1. 初始化当前节点状态
    state := initState()
    
    // 2. 递归处理子节点
    for _, child := range children[u] {
        childState := dfs(child)
        
        // 3. 合并子节点状态
        state = merge(state, childState)
    }
    
    // 4. 根据子节点状态更新当前节点
    state = update(state, u)
    
    return state
}
```

#### 子节点合并策略

| 问题类型 | 合并方式 | 示例 |
|---------|---------|------|
| 独立子问题 | 直接求和 | `sum(dp[child])` |
| 有约束条件 | 背包DP | 本题的预算分配 |
| 互斥选择 | 取最大/最小 | 只能选一个子节点 |

#### 经典树形DP问题

**1. 打家劫舍 III (House Robber III)**
```go
// 状态: dp[u][0/1] = 不选/选节点u的最大价值
dp[u][0] = sum(max(dp[child][0], dp[child][1]))
dp[u][1] = value[u] + sum(dp[child][0])
```

**2. 树的直径**
```go
// 状态: dp[u] = 以u为端点的最长路径
diameter = max(diameter, dp[child1] + dp[child2] + 2)
dp[u] = max(dp[child]) + 1
```

**3. 本题（树形DP + 背包）**
```go
// 状态: dfs(u) -> {dp0[], dp1[], size}
// 一次返回所有预算的最优解
```

### 8.3 本题算法思维链条

```
识别树形结构
    ↓
考虑树形DP
    ↓
发现预算约束 → 结合背包问题
    ↓
父子节点相互影响 → 状态需包含父节点信息
    ↓
多个子节点 → 需要分组背包分配预算
    ↓
性能瓶颈 → 批量计算所有预算 + size剪枝
    ↓
最终解法：树形DP + 分组背包 + 批量计算优化
```

## 九、关键技巧总结

### 1. 识别问题类型
- **树形结构** → 树形DP
- **预算限制** → 背包问题
- **组合优化** → 分组背包

### 2. 状态设计优化
```go
// ❌ 朴素设计：dfs(u, budget, parentBought) -> int
// ✅ 优化设计：dfs(u) -> result{dp0[], dp1[], size}
```
**关键**：一次返回所有预算的结果，空间换时间

### 3. 剪枝技巧
- 用 `size` 限制背包搜索空间
- 子树最多花费 `size`，超过此预算的遍历无意义
- 在大数据下效果显著

### 4. 分组背包注意事项
- **必须倒序遍历**预算，避免重复使用子节点
- 每个子节点是一个"组"
- 在组内选择最优分配方案

## 十、扩展思考

1. **如果允许卖空？** 需要增加状态维度表示持仓情况
2. **如果折扣策略更复杂？** 可以将折扣率作为参数传递
3. **如果是多叉树且深度很大？** 考虑记忆化或启发式剪枝

## 十一、参考资料

- [树形DP经典问题 - OI Wiki](https://oi-wiki.org/dp/tree/)
- [背包问题九讲](https://github.com/tianyicui/pack)
- [LeetCode 3562 官方题解](https://leetcode.cn/problems/maximum-profit-from-trading-stocks-with-discounts/solutions/)

---

**总结**：本题是树形DP与分组背包的经典结合，关键优化在于**批量计算所有预算的结果**和**size 剪枝**。这种"空间换时间"的思路将时间复杂度从 O(n×budget³) 优化到 O(n×budget²)，使得大规模测试用例能够通过。掌握这类优化技巧，有助于解决更复杂的树上优化问题。
