# 📝 LeetCode 2054：两个最好的不重叠活动 —— 从暴力到优化的完整思路

> **适合人群**：有一定算法基础，想深入理解区间调度 + 动态规划的程序员
> **难度**：中等
> **核心知识点**：排序、动态规划、二分查找、前缀最大值优化

---

## 🎯 问题描述

给你一个二维数组 `events`，其中 `events[i] = [startTime, endTime, value]`。

- 第 i 个活动开始于 `startTime[i]`，结束于 `endTime[i]`
- 如果你参加这个活动，可以获得价值 `value[i]`
- 你**最多可以参加 2 个不重叠的活动**，求最大价值

**不重叠定义**：如果活动 A 结束于时间 t，那么活动 B 必须在 `t+1` 或之后开始。

---

## 📌 示例

### 示例 1
```
输入：events = [[1,3,2],[4,5,2],[2,4,3]]
输出：4
解释：选择活动 0 和活动 1，价值 = 2 + 2 = 4
```

### 示例 2
```
输入：events = [[1,3,2],[4,5,2],[1,5,5]]
输出：5
解释：选择活动 2，价值 = 5
```

### 示例 3
```
输入：events = [[1,5,3],[1,5,1],[6,6,5]]
输出：8
解释：选择活动 0 和活动 2，价值 = 3 + 5 = 8
```

---

## 🤔 思路分析

### 第一反应：背包问题？

很多人（包括我）第一眼会想到 **0-1 背包**：
- 物品 = 活动
- 重量 = 时间跨度
- 价值 = value
- 容量 = 总时间范围

**但这是错的！** ❌

**原因**：
- 背包问题的约束是"总重量不超过容量"
- 这道题的约束是"时间不能重叠"
- **时间重叠 ≠ 重量累加**

举个反例：
```
活动 A: [1, 5, 100]  (时间跨度 4)
活动 B: [2, 3, 50]   (时间跨度 1)
```
如果用背包思路，A 和 B 的"总重量"是 5，看起来可以装下。
但实际上它们**时间重叠**，不能同时选！

---

### 正确思路：区间调度问题

这是一个**区间调度问题**的变种：
- 经典的区间调度：选尽可能多的不重叠活动（贪心）
- 这道题：最多选 2 个，求最大价值（动态规划）

---

## 🔍 解法一：暴力枚举（O(n²)）

### 核心思想
既然最多选 2 个，那就**穷举所有可能的活动对**：
1. 只选 1 个活动 i → 价值 = `value[i]`
2. 选 2 个活动 (i, j) → 如果不冲突，价值 = `value[i] + value[j]`

### 代码实现
```python
def maxTwoEvents(events):
    n = len(events)
    max_value = 0
  
    # 1. 只选一个活动
    for i in range(n):
        max_value = max(max_value, events[i][2])
  
    # 2. 选两个活动
    for i in range(n):
        for j in range(i + 1, n):
            # 判断是否不重叠
            if events[i][1] < events[j][0] or events[j][1] < events[i][0]:
                max_value = max(max_value, events[i][2] + events[j][2])
  
    return max_value
```

### 复杂度分析
- **时间复杂度**：O(n²)
- **空间复杂度**：O(1)

### 能否通过？
- 对于 n ≤ 1000，可以通过
- 对于 n ≤ 10⁵（LeetCode 的数据范围），会**超时** ❌

---

## 🚀 解法二：排序 + 动态规划 + 二分查找（O(n log n)）

### 核心思想

#### 1️⃣ 先排序
按 `endTime` **升序**排序，这样：
- 处理到活动 i 时，前面所有活动的结束时间都 ≤ `events[i][1]`
- 方便用二分查找找到"最后一个不冲突的活动"

#### 2️⃣ 动态规划状态定义
```
dp[i] = 选择第 i 个活动时，能获得的最大价值
```

**状态转移**：
- **情况 1**：只选活动 i → `dp[i] = value[i]`
- **情况 2**：选活动 i + 前面某个不冲突的活动 j → `dp[i] = value[i] + dp[j]`

其中，j 是满足 `events[j].endTime < events[i].startTime` 的**最后一个活动**。

#### 3️⃣ 二分查找优化
用二分查找快速找到 j 的位置：
```python
def binary_search(events, i):
    """找到最后一个 endTime < events[i].startTime 的活动"""
    left, right = 0, i - 1
    result = -1
  
    while left <= right:
        mid = (left + right) // 2
        if events[mid][1] < events[i][0]:
            result = mid
            left = mid + 1
        else:
            right = mid - 1
  
    return result
```

#### 4️⃣ 前缀最大值优化
如果直接用 `dp[i] = value[i] + max(dp[0:j])`，每次需要 O(n) 扫描。

**优化**：维护一个 `max_dp[i]`，表示前 i 个活动中 dp 的最大值：
```python
max_dp[i] = max(max_dp[i-1], dp[i])
```

这样查询变成 O(1)！

---

### 完整代码

```python
def maxTwoEvents(events):
    # 1. 按结束时间排序
    events.sort(key=lambda x: x[1])
    n = len(events)
  
    # 2. 初始化 DP
    dp = [0] * n      # dp[i] = 选第 i 个活动的最大价值
    max_dp = [0] * n  # max_dp[i] = 前 i 个活动中 dp 的最大值
  
    # 3. 动态规划
    for i in range(n):
        # 情况 1：只选活动 i
        dp[i] = events[i][2]
      
        # 情况 2：选活动 i + 前面某个不冲突的活动
        j = binary_search(events, i)
        if j != -1:
            dp[i] = max(dp[i], events[i][2] + max_dp[j])
      
        # 更新前缀最大值
        max_dp[i] = max(max_dp[i-1] if i > 0 else 0, dp[i])
  
    # 4. 返回最大值
    return max_dp[n-1]


def binary_search(events, i):
    """找到最后一个 endTime < events[i].startTime 的活动"""
    left, right = 0, i - 1
    result = -1
  
    while left <= right:
        mid = (left + right) // 2
        if events[mid][1] < events[i][0]:
            result = mid
            left = mid + 1
        else:
            right = mid - 1
  
    return result
```

---

### 复杂度分析

| 操作 | 时间复杂度 |
|------|-----------|
| 排序 | O(n log n) |
| 遍历每个活动 | O(n) |
| 每次二分查找 | O(log n) |
| **总计** | **O(n log n)** |

**空间复杂度**：O(n)（dp 和 max_dp 数组）

---

## 🧩 算法流程图

```
输入: events = [[1,3,2],[4,5,2],[2,4,3]]
         ↓
排序 (按 endTime 升序)
         ↓
events = [[1,3,2],[2,4,3],[4,5,2]]
         ↓
遍历每个活动 i:
  - 只选 i: dp[i] = value[i]
  - 二分查找: 找最后一个不冲突的 j
  - 选 i+j: dp[i] = value[i] + max_dp[j]
  - 更新: max_dp[i] = max(max_dp[i-1], dp[i])
         ↓
返回 max_dp[n-1]
```

---

## 🎨 时间轴可视化

```
示例: events = [[1,3,2],[4,5,2],[2,4,3]]

排序后:
活动 0: [1----3] value=2
活动 1:   [2---4] value=3
活动 2:       [4----5] value=2

时间轴:
0  1  2  3  4  5  6
   [--0--]
      [-1--]
         [--2--]

DP 过程:
i=0: dp[0]=2, max_dp[0]=2  (只选活动0)
i=1: dp[1]=3, max_dp[1]=3  (只选活动1，价值更高)
i=2: 
  - 只选2: dp[2]=2
  - 选0+2: events[0].endTime=3 < events[2].startTime=4 ✅
           dp[2] = 2 + max_dp[0] = 2 + 2 = 4
  - max_dp[2] = max(3, 4) = 4

答案: 4
```

---

## 🔑 关键技巧总结

### 1. 为什么按 endTime 排序？
- 处理活动 i 时，前面所有活动都已"结束"
- 方便用二分查找找到"最后一个不冲突的活动"
- 符合"从前往后"的 DP 思路

### 2. 二分查找的边界
```python
# 我们要找：最后一个满足 events[j].endTime < events[i].startTime 的 j
# 即：events[j][1] < events[i][0]

# 标准的 lower_bound 变体
if events[mid][1] < events[i][0]:
    result = mid
    left = mid + 1  # 继续往右找，看有没有更大的
else:
    right = mid - 1
```

### 3. 前缀最大值的妙用
- 不用每次 O(n) 扫描 `max(dp[0:j])`
- 维护 `max_dp[i] = max(max_dp[i-1], dp[i])`
- 查询变成 O(1)

---

## 🧪 测试用例

```python
# 测试 1：基础情况
events = [[1,3,2],[4,5,2],[2,4,3]]
assert maxTwoEvents(events) == 4

# 测试 2：只选一个更优
events = [[1,3,2],[4,5,2],[1,5,5]]
assert maxTwoEvents(events) == 5

# 测试 3：两个不相邻
events = [[1,5,3],[1,5,1],[6,6,5]]
assert maxTwoEvents(events) == 8

# 测试 4：全部重叠
events = [[1,5,10],[2,4,20],[3,6,15]]
assert maxTwoEvents(events) == 20

# 测试 5：全部不重叠
events = [[1,2,1],[3,4,2],[5,6,3]]
assert maxTwoEvents(events) == 5  # 选后两个
```

---

## 💡 扩展思考

### 如果改成"最多选 k 个活动"？
- 状态定义：`dp[i][k] = 前 i 个活动中，选恰好 k 个的最大价值`
- 时间复杂度：O(n² · k)（可能需要线段树优化）

### 如果活动价值可以叠加多次？
- 变成"加权区间调度问题"（Weighted Interval Scheduling）
- 经典 DP 问题，贪心不再适用

---

## 📚 相关题目

- [LeetCode 435. 无重叠区间](https://leetcode.com/problems/non-overlapping-intervals/)（贪心）
- [LeetCode 452. 用最少数量的箭引爆气球](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)（贪心）
- [LeetCode 1235. 规划兼职工作](https://leetcode.com/problems/maximum-profit-in-job-scheduling/)（本题的 k=∞ 版本）

---

## 🎓 总结

| 方法 | 时间复杂度 | 空间复杂度 | 适用范围 |
|------|-----------|-----------|---------|
| 暴力枚举 | O(n²) | O(1) | n ≤ 1000 |
| DP + 二分 + 前缀优化 | O(n log n) | O(n) | n ≤ 10⁵ ✅ |

**核心思想**：
1. **排序**减少搜索空间
2. **动态规划**避免重复计算
3. **二分查找**加速查询
4. **前缀最大值**优化状态转移

**程序员视角**：
- 排序 = 预处理，让数据变得"有序可查"
- DP = 缓存，避免重复计算
- 二分 = 索引优化，把 O(n) 查询降到 O(log n)
- 前缀最大值 = 空间换时间，类似数据库的物化视图

---

## 🙏 写在最后

这道题是一个很好的**算法组合拳**练习：
- 不是单纯的贪心或 DP
- 需要**排序 + DP + 二分 + 前缀优化**四个技巧配合
- 体现了"分而治之"的思想

如果你能独立写出这道题，说明你已经具备了：
✅ 问题抽象能力（识别出这是区间调度问题）
✅ 算法设计能力（选择合适的数据结构和算法）
✅ 优化思维（从 O(n²) 优化到 O(n log n)）

**继续加油！** 💪

---

> **如果这篇文章对你有帮助，欢迎点赞、收藏、分享！**
> **有问题欢迎在评论区讨论～** 😊