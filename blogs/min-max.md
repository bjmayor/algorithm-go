# 最大化最小值 / 最小化最大值 - 算法套路总结

## 一、什么是"最大化最小值"问题？

在算法题中，经常会遇到这样的描述：
- "让所有工人的最大工作量尽可能小"
- "让所有城市的最小电量尽可能大"  
- "让木材切割后的最小长度尽可能大"
- "让学生分配到的最大任务数尽可能小"

这类问题有个共同特点：**求在某种约束条件下，某个极值指标的最优值**。

乍一看这类问题很难直接求解，但有一个巧妙的思路：**把"求最优值"转换为"判断某个值是否可行"**。

---

## 二、核心思想：二分答案法

### 2.1 为什么要二分答案？

**直接思路的困境：**
假设要求"分配任务后，学生们的最大任务量的最小值"，如果直接枚举所有可能的分配方案，复杂度会是指数级的，根本无法接受。

**转换思路：**
我们换个角度问问题：
- 原问题："最优答案是多少？"（不知道怎么求）
- 新问题："答案能否达到 X？"（相对容易判断）

如果能高效判断"答案是否能达到 X"，那我们就可以用**二分搜索**在所有可能的答案中找到最优的那个！

### 2.2 为什么二分法有效？关键在于单调性

二分答案法之所以可行，是因为这类问题有一个关键性质：**单调性**。

以"最大化最小值"为例：
- 如果"让所有城市电量都达到 10"**可以实现**，那么"让所有城市电量都达到 9、8、7..."肯定也**可以实现**（要求更低了）
- 如果"让所有城市电量都达到 10"**无法实现**，那么"让所有城市电量都达到 11、12、13..."肯定也**无法实现**（要求更高了）

这种单调性意味着：**存在一个临界值 M，小于等于 M 的都可以实现，大于 M 的都无法实现**。

```
答案值:  0   1   2   3   4   5   6   7   8   9   10  11  12
是否可行: ✓   ✓   ✓   ✓   ✓   ✓   ✓   ✓   ✗   ✗   ✗   ✗   ✗
                                      ↑
                                   我们要找的最优答案
```

这正是二分搜索的应用场景！我们要找的就是那个"从可行到不可行"的临界点。

---

## 三、解题框架（四步走）

### 第一步：识别问题类型

看到以下关键词就要想到二分答案法：
- "最大化最小值"
- "最小化最大值"  
- "在...约束下，求...的最优值"
- "让...尽可能大/小"

### 第二步：确定二分搜索范围

需要确定答案的可能范围 `[left, right]`：

- **left（最小可能值）**：通常是理论最小值，比如 0，或者初始状态下已有的最小值
- **right（最大可能值）**：通常是理论最大值，比如所有资源之和，或者题目给定的上限

### 第三步：设计 check(mid) 函数

这是**最关键**的一步！

`check(mid)` 函数的作用：**判断答案能否达到 mid**。

- 返回 `true`：说明 mid 这个答案可以实现，真正的最优答案 ≥ mid
- 返回 `false`：说明 mid 这个答案无法实现，真正的最优答案 < mid

设计 check 函数时，通常会用到**贪心、滑动窗口、差分数组**等技巧。

### 第四步：二分搜索主循环

```python
def solve():
    left, right = 确定搜索范围
    
    while left < right:
        mid = (left + right + 1) // 2  # 对于"最大化最小值"问题
        
        if check(mid):  # mid 可行
            left = mid  # 尝试更大的值
        else:  # mid 不可行
            right = mid - 1  # 尝试更小的值
    
    return left  # 最终 left == right，就是答案
```

**注意：**
- "最大化最小值"问题：`mid = (left + right + 1) // 2`，可行时 `left = mid`
- "最小化最大值"问题：`mid = (left + right) // 2`，可行时 `right = mid`

---

## 四、完整案例：分木材问题

### 4.1 问题描述

你有 n 根木材，第 i 根的长度是 lengths[i]。你需要切割这些木材得到至少 k 根木材。

**求：切割后木材的最大可能长度是多少？**

**示例：**
```
输入：lengths = [5, 9, 7], k = 3
输出：5

解释：
- 长度为 5 的木材：保持不切，得到 1 根长度为 5 的木材
- 长度为 9 的木材：切成 5+4，得到 1 根长度为 5 的木材（4 舍弃）
- 长度为 7 的木材：切成 5+2，得到 1 根长度为 5 的木材（2 舍弃）
共得到 3 根长度为 5 的木材，满足要求。

如果目标长度是 6：
- 长度为 5 的木材：无法切出长度为 6 的木材，得到 0 根
- 长度为 9 的木材：切成 6+3，得到 1 根
- 长度为 7 的木材：切成 6+1，得到 1 根
只能得到 2 根，不满足要求。
```

### 4.2 问题分析

这是一个典型的"最大化最小值"问题：要让切割后木材的长度（最小的那个标准）尽可能大。

**思路转换：**
- 原问题："切割后的最大长度是多少？" → 不知道怎么直接求
- 新问题："能否切出至少 k 根长度为 X 的木材？" → 可以遍历检查

### 4.3 解题步骤

**第一步：确定搜索范围**
```python
left = 1  # 最小长度至少为 1
right = max(lengths)  # 最大长度不会超过最长的木材
```

**第二步：设计 check(mid) 函数**

判断"能否切出至少 k 根长度为 mid 的木材"：

```python
def check(mid):
    count = 0  # 统计能切出多少根长度为 mid 的木材
    
    for length in lengths:
        count += length // mid  # 每根木材能切出几根
    
    return count >= k  # 至少 k 根就算成功
```

**为什么这样可行？**
- 每根木材 `length` 可以切出 `length // mid` 根长度为 `mid` 的木材
- 把所有木材能切出的数量加起来，如果 ≥ k，说明目标长度 `mid` 是可行的
- 这是一个贪心策略：每根木材尽可能多地切出目标长度，不浪费

**第三步：二分搜索主循环**

```python
def maxLength(lengths, k):
    left, right = 1, max(lengths)
    
    while left < right:
        mid = (left + right + 1) // 2  # 向上取整，避免死循环
        
        if check(mid):  # 可以切出至少 k 根长度为 mid 的木材
            left = mid  # 尝试更长的长度
        else:  # 无法切出足够的木材
            right = mid - 1  # 降低长度要求
    
    return left
```

### 4.4 完整代码实现

```python
def cutWood(lengths, k):
    """
    lengths: 木材长度数组
    k: 需要得到的木材根数
    返回：切割后木材的最大可能长度
    """
    def check(target_length):
        """检查能否切出至少 k 根长度为 target_length 的木材"""
        count = 0
        for length in lengths:
            count += length // target_length
        return count >= k
    
    # 确定搜索范围
    left, right = 1, max(lengths)
    
    # 二分搜索
    while left < right:
        mid = (left + right + 1) // 2  # 向上取整
        
        if check(mid):
            left = mid  # mid 可行，尝试更大的值
        else:
            right = mid - 1  # mid 不可行，尝试更小的值
    
    return left

# 测试
lengths = [5, 9, 7]
k = 3
print(cutWood(lengths, k))  # 输出: 5
```

### 4.5 执行过程图解

```
初始：lengths = [5, 9, 7], k = 3
搜索范围：left = 1, right = 9

第一轮：
  mid = (1 + 9 + 1) // 2 = 5
  check(5): 5//5=1, 9//5=1, 7//5=1 → 共 3 根 ✓
  结论：5 可行，尝试更大的 → left = 5

第二轮：
  left = 5, right = 9
  mid = (5 + 9 + 1) // 2 = 7
  check(7): 5//7=0, 9//7=1, 7//7=1 → 共 2 根 ✗
  结论：7 不可行，降低要求 → right = 6

第三轮：
  left = 5, right = 6
  mid = (5 + 6 + 1) // 2 = 6
  check(6): 5//6=0, 9//6=1, 7//6=1 → 共 2 根 ✗
  结论：6 不可行，降低要求 → right = 5

第四轮：
  left = 5, right = 5 → 循环结束
  
答案：5
```

---

## 五、案例 2：分配任务（最小化最大值）

### 5.1 问题描述

有 n 个任务，每个任务的工作量是 tasks[i]。现在要把这些任务分配给 k 个工人，每个工人可以完成多个任务，但必须按顺序完成（不能跳过）。

**求：如何分配，使得工人中的最大工作量尽可能小？**

**示例：**
```
输入：tasks = [1, 2, 3, 4, 5], k = 3
输出：6

解释：
工人1: [1, 2, 3] → 工作量 6
工人2: [4]       → 工作量 4  
工人3: [5]       → 工作量 5
最大工作量是 6

如果分成：
工人1: [1, 2] → 4
工人2: [3, 4] → 7  (最大)
工人3: [5]    → 5
最大工作量是 7，不如第一种分配
```

### 5.2 解题步骤

**第一步：确定搜索范围**
```python
left = max(tasks)  # 至少要能完成最大的单个任务
right = sum(tasks)  # 最多是一个人完成所有任务
```

**第二步：设计 check(mid) 函数**

判断"能否用 k 个工人，让每个人的工作量都不超过 mid"：

```python
def check(max_workload):
    workers_needed = 1  # 至少需要 1 个工人
    current_workload = 0  # 当前工人的工作量
    
    for task in tasks:
        if current_workload + task <= max_workload:
            # 当前工人可以继续做这个任务
            current_workload += task
        else:
            # 需要新工人
            workers_needed += 1
            current_workload = task
    
    return workers_needed <= k  # 需要的工人数不超过 k
```

**为什么这样可行？**
- 贪心策略：让每个工人尽可能多做任务，直到无法再做
- 如果需要的工人数 ≤ k，说明这个 max_workload 是可行的
- 如果需要的工人数 > k，说明 max_workload 太小了，需要增大

**第三步：二分搜索**

```python
def minimizeMaxWorkload(tasks, k):
    left, right = max(tasks), sum(tasks)
    
    while left < right:
        mid = (left + right) // 2  # 向下取整
        
        if check(mid):
            right = mid  # 可行，尝试更小的值
        else:
            left = mid + 1  # 不可行，需要更大的值
    
    return left
```

**注意：** 这里是"最小化最大值"问题，所以：
- 使用 `mid = (left + right) // 2`（向下取整）
- 可行时 `right = mid`（尝试更小）

### 5.3 完整代码

```python
def minimizeMaxWorkload(tasks, k):
    def check(max_workload):
        """检查能否用 k 个工人完成，每人工作量不超过 max_workload"""
        workers_needed = 1
        current_workload = 0
        
        for task in tasks:
            if current_workload + task <= max_workload:
                current_workload += task
            else:
                workers_needed += 1
                current_workload = task
                
                if workers_needed > k:
                    return False
        
        return True
    
    left, right = max(tasks), sum(tasks)
    
    while left < right:
        mid = (left + right) // 2
        
        if check(mid):
            right = mid
        else:
            left = mid + 1
    
    return left

# 测试
tasks = [1, 2, 3, 4, 5]
k = 3
print(minimizeMaxWorkload(tasks, k))  # 输出: 6
```

---

## 六、案例 3：LeetCode 2528 最大化城市的最小电量

### 6.1 问题描述

给你一个下标从 0 开始、长度为 n 的整数数组 `stations`，其中 `stations[i]` 表示第 i 座城市的电站数量。

每个电站的供电范围是固定的。也就是说，每个电站可以给它所在的城市以及距离它不超过 r 的城市供电。

你可以新建最多 k 个电站。每个新建电站的供电范围和已有电站相同。

返回所有城市电量的最小值的最大值。

**示例：**
```
输入：stations = [1, 2, 4, 5, 0], r = 1, k = 2
输出：5

解释：
初始状态下每个城市的电量：
- 城市 0: stations[0] + stations[1] = 1 + 2 = 3
- 城市 1: stations[0] + stations[1] + stations[2] = 1 + 2 + 4 = 7
- 城市 2: stations[1] + stations[2] + stations[3] = 2 + 4 + 5 = 11
- 城市 3: stations[2] + stations[3] + stations[4] = 4 + 5 + 0 = 9
- 城市 4: stations[3] + stations[4] = 5 + 0 = 5

最小电量是 3（城市 0）。

如果在城市 1 新建 2 个电站：
- 城市 0: 3 + 2 = 5
- 城市 1: 7 + 2 = 9
- 城市 2: 11 + 2 = 13
- 城市 3: 9
- 城市 4: 5

所有城市电量最小值变成 5。
```

### 6.2 问题分析

这是一个典型的"最大化最小值"问题，而且难度更高，因为：
1. 电站有覆盖范围，不是简单的一对一
2. 需要计算每个城市的初始电量（区间和）
3. check 函数需要模拟新建电站的过程

**思路转换：**
- 原问题："所有城市电量的最小值最大是多少？" → 难以直接求解
- 新问题："能否通过新建最多 k 个电站，让所有城市电量都至少达到 X？" → 可以贪心检查

### 6.3 解题步骤

**第一步：确定搜索范围**
```python
# 最小值：初始状态下的最小电量
left = 0
# 最大值：所有电站之和 + 新建的 k 个电站
right = sum(stations) + k
```

**第二步：计算每个城市的初始电量**

由于电站的覆盖范围是 r，城市 i 的电量 = stations[i-r] + ... + stations[i+r]。
这是一个典型的滑动窗口问题：

```python
def get_initial_power(stations, r):
    n = len(stations)
    power = [0] * n
    
    # 使用滑动窗口计算每个城市的初始电量
    window_sum = sum(stations[max(0, i - r):min(n, i + r + 1)] 
                     for i in range(1))
    
    for i in range(n):
        # 计算城市 i 的覆盖范围 [i-r, i+r]
        left_idx = max(0, i - r)
        right_idx = min(n - 1, i + r)
        power[i] = sum(stations[left_idx:right_idx + 1])
    
    return power
```

**第三步：设计 check(mid) 函数**

判断"能否用最多 k 个电站，让所有城市电量都至少达到 mid"：

核心思想：
1. 从左到右遍历每个城市
2. 如果当前城市电量 < mid，需要新建电站
3. **贪心策略**：在能覆盖当前城市的最右侧位置建电站（这样能同时帮助后面的城市）
4. 使用差分数组高效处理区间增加

```python
def check(mid):
    # 复制一份电站数组，用于模拟新建
    add = [0] * n  # 新建的电站数量
    power = 0  # 当前位置的累计电量
    used = 0  # 已使用的新建电站数
    
    # 初始化第一个城市的电量
    for j in range(min(n, r + 1)):
        power += stations[j] + add[j]
    
    for i in range(n):
        # 更新滑动窗口
        if i > 0:
            # 移出左边界
            left_out = i - r - 1
            if left_out >= 0:
                power -= stations[left_out] + add[left_out]
            # 加入右边界
            right_in = i + r
            if right_in < n:
                power += stations[right_in] + add[right_in]
        
        # 如果当前城市电量不足
        if power < mid:
            need = mid - power
            # 在最右侧能覆盖当前城市的位置建电站
            build_pos = min(n - 1, i + r)
            add[build_pos] += need
            power += need
            used += need
            
            if used > k:
                return False
    
    return True
```

**为什么在最右侧建电站？**
- 电站建在位置 j，能覆盖 [j-r, j+r] 范围
- 如果城市 i 电量不足，可以在 [i-r, i+r] 范围内任意位置建电站
- 选择 min(n-1, i+r)（最右侧）能让这个电站同时帮助 i 后面的城市
- 这是贪心策略：让每个电站的价值最大化

### 6.4 完整代码实现

```python
def maxPower(stations, r, k):
    n = len(stations)
    
    def check(min_power):
        """检查能否让所有城市电量都达到 min_power"""
        add = [0] * n  # 新建的电站
        power = 0  # 当前滑动窗口的电量和
        used = 0  # 已使用的电站数
        
        # 初始化城市 0 的电量
        for j in range(min(n, r + 1)):
            power += stations[j]
        
        for i in range(n):
            # 维护滑动窗口 [i-r, i+r]
            if i > 0:
                # 移除左边界
                if i - r - 1 >= 0:
                    power -= stations[i - r - 1] + add[i - r - 1]
                # 添加右边界
                if i + r < n:
                    power += stations[i + r] + add[i + r]
            
            # 当前城市电量不足
            if power < min_power:
                need = min_power - power
                # 在最右侧能覆盖当前城市的位置建电站
                build_pos = min(n - 1, i + r)
                add[build_pos] += need
                power += need
                used += need
                
                if used > k:
                    return False
        
        return True
    
    # 计算初始最小电量作为 left
    power = [0] * n
    for i in range(n):
        for j in range(max(0, i - r), min(n, i + r + 1)):
            power[i] += stations[j]
    
    left = min(power)
    right = sum(stations) + k
    
    # 二分搜索
    while left < right:
        mid = (left + right + 1) // 2
        
        if check(mid):
            left = mid
        else:
            right = mid - 1
    
    return left

# 测试
stations = [1, 2, 4, 5, 0]
r = 1
k = 2
print(maxPower(stations, r, k))  # 输出: 5
```

### 6.5 Go 语言实现

```go
func maxPower(stations []int, r int, k int) int64 {
    n := len(stations)
    
    // 计算初始电量
    power := make([]int64, n)
    for i := 0; i < n; i++ {
        for j := max(0, i-r); j < min(n, i+r+1); j++ {
            power[i] += int64(stations[j])
        }
    }
    
    // 找到初始最小电量
    left := power[0]
    for _, p := range power {
        if p < left {
            left = p
        }
    }
    
    // 计算最大可能电量
    right := int64(0)
    for _, s := range stations {
        right += int64(s)
    }
    right += int64(k)
    
    // 二分搜索
    for left < right {
        mid := (left + right + 1) / 2
        if check(stations, r, k, mid) {
            left = mid
        } else {
            right = mid - 1
        }
    }
    
    return left
}

func check(stations []int, r int, k int, minPower int64) bool {
    n := len(stations)
    add := make([]int64, n)
    power := int64(0)
    used := int64(0)
    
    // 初始化城市 0 的电量
    for j := 0; j < min(n, r+1); j++ {
        power += int64(stations[j])
    }
    
    for i := 0; i < n; i++ {
        // 维护滑动窗口
        if i > 0 {
            if i-r-1 >= 0 {
                power -= int64(stations[i-r-1]) + add[i-r-1]
            }
            if i+r < n {
                power += int64(stations[i+r]) + add[i+r]
            }
        }
        
        // 电量不足，需要新建
        if power < minPower {
            need := minPower - power
            buildPos := min(n-1, i+r)
            add[buildPos] += need
            power += need
            used += need
            
            if used > int64(k) {
                return false
            }
        }
    }
    
    return true
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

### 6.6 关键技巧总结

这道题综合了多个技巧：

1. **滑动窗口**：高效计算每个城市的电量（区间和）
2. **贪心策略**：电站建在最右侧，让效益最大化
3. **差分思想**：虽然代码中直接维护 add 数组，但本质是差分的思想
4. **二分答案**：将"求最优值"转换为"判断可行性"

**时间复杂度：** O(n * log(sum + k))
- 二分搜索：O(log(sum + k))
- 每次 check：O(n)

---

## 七、常见问题类型汇总

| 问题类型 | 二分目标 | check 函数判断 | 典型题目 |
|---------|---------|---------------|---------|
| 分配问题 | 最大工作量的最小值 | 能否用 k 个人完成 | 分配任务、分配书籍 |
| 切割问题 | 最大长度 | 能否切出 k 根 | 切割木材、绳子 |
| 距离问题 | 最小距离的最大值 | 能否满足距离要求 | 放置磁铁、安排座位 |
| 资源分配 | 最小资源的最大值 | 能否满足所有需求 | 供电问题、供水问题 |
| 速度问题 | 最小速度 | 能否在规定时间完成 | 运输货物、完成任务 |

---

## 七、解题技巧总结

### 7.1 如何确定 mid 的计算方式？

**最大化最小值问题：**
```python
mid = (left + right + 1) // 2  # 向上取整
if check(mid):
    left = mid  # 可行，尝试更大
else:
    right = mid - 1  # 不可行，降低要求
```

**最小化最大值问题：**
```python
mid = (left + right) // 2  # 向下取整
if check(mid):
    right = mid  # 可行，尝试更小
else:
    left = mid + 1  # 不可行，提高限制
```

### 7.2 check 函数常用技巧

1. **贪心策略**：尽可能充分利用每个资源
2. **滑动窗口**：处理区间覆盖问题
3. **差分数组**：处理区间修改问题
4. **前缀和**：快速计算区间和

### 7.3 时间复杂度分析

- 二分次数：O(log(right - left))
- 每次 check：O(n)
- 总复杂度：O(n * log(right - left))

通常比暴力枚举的 O(n²) 甚至 O(2^n) 快得多！

---

## 八、练习建议

掌握这个套路需要三步：
1. **识别模式**：看到题目能立刻想到"这是二分答案"
2. **设计 check**：这是难点，需要结合贪心、滑动窗口等技巧
3. **注意细节**：搞清楚是向上还是向下取整，left/right 如何更新

建议从简单题目开始练习：
- LeetCode 875（爱吃香蕉的珂珂）
- LeetCode 1011（在 D 天内送达包裹的能力）
- LeetCode 410（分割数组的最大值）
- LeetCode 1482（制作 m 束花所需的最少天数）

祝你早日掌握这个强大的算法套路！
