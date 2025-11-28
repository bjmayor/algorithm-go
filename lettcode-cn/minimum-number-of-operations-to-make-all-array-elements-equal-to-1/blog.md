# LeetCode 2654. 使数组所有元素变成 1 的最少操作次数 - GCD 思维题详解

## 问题分析

这道题乍一看可能会让人想到动态规划或贪心算法，但实际上它更像是一道数学思维题。关键在于理解最大公约数（GCD）的特性和传播规律。

## 核心洞察

### GCD 的关键性质

1. **1 的特殊地位**：`gcd(1, x) = 1` 对任何正整数 x 都成立
2. **传播性**：如果有一个 1，它可以通过操作"感染"相邻的元素
3. **子数组特性**：如果一个子数组的 gcd 为 1，就可以通过操作使整个子数组变成 1

### GCD 的结合律性质

**GCD 结合律**：`gcd(a, b, c) = gcd(gcd(a, b), c)`

这个性质是理解本题的关键！它告诉我们：
- 即使 `gcd(a, b) ≠ 1` 和 `gcd(b, c) ≠ 1`
- 仍然可能 `gcd(a, b, c) = 1`

**示例**：`[6, 10, 15]`
- `gcd(6, 10) = 2`，`gcd(10, 15) = 5`
- 但是 `gcd(6, 10, 15) = gcd(gcd(6, 10), 15) = gcd(2, 15) = 1`

这就是为什么需要检查整个数组的 gcd，而不仅仅是相邻元素！

### 欧几里得算法

计算 GCD 的标准方法是欧几里得算法：

**算法原理**：`gcd(a, b) = gcd(b, a % b)`，直到 `b = 0`

**Go 实现**：
```go
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
```

**示例计算**：
- `gcd(48, 18) = gcd(18, 48 % 18) = gcd(18, 12)`
- `gcd(18, 12) = gcd(12, 18 % 12) = gcd(12, 6)`
- `gcd(12, 6) = gcd(6, 12 % 6) = gcd(6, 0)`
- `gcd(6, 0) = 6`

### 操作的本质理解

题目中的操作实际上是：选择相邻的两个数，用它们的 gcd 替换其中一个。
这相当于我们可以在数组中"传播" gcd 值。

## 解题思路详解

### 情况一：数组中已有 1

这是最简单的情况。如果数组中有 `count` 个 1，那么：

- 每个非 1 元素只需要 1 次操作就可以变成 1
- 总操作次数 = `数组长度 - count`

**示例**：`[1, 3, 6, 1, 4]`
- 有 2 个 1
- 需要将 3 个非 1 元素变成 1
- 操作次数 = 5 - 2 = 3

### 情况二：数组中没有 1，但整体 gcd = 1

如果数组中没有 1，但整个数组的 gcd 为 1，说明我们**可以**通过操作得到 1。

关键思路：找到最短的子数组，使其 gcd 为 1。

**为什么是最短子数组？**
- 长度为 k 的子数组，需要 k-1 次操作来得到第一个 1
- 得到 1 后，需要 n-1 次操作传播到整个数组
- 总次数 = (k-1) + (n-1) = n + k - 2
- 要使总次数最小，就要使 k 最小

### 情况三：不可能的情况

如果整个数组的 gcd > 1，那么无论如何操作都不可能得到 1，直接返回 -1。

## 算法实现

### 寻找最短 gcd=1 子数组的方法

```go
func findShortestSubarray(nums []int) int {
    n := len(nums)
    minLen := n + 1 // 初始化为一个不可能的值
    
    // 枚举所有子数组
    for i := 0; i < n; i++ {
        currentGcd := nums[i]
        if currentGcd == 1 {
            return 1 // 找到长度为1的子数组
        }
        
        for j := i + 1; j < n; j++ {
            currentGcd = gcd(currentGcd, nums[j])
            if currentGcd == 1 {
                minLen = min(minLen, j - i + 1)
                break // 找到了从这个位置开始的最短子数组
            }
        }
    }
    
    return minLen
}
```

### 完整解题步骤

1. 统计数组中 1 的个数
2. 如果有 1，直接返回 `n - count`
3. 如果没有 1，检查整个数组的 gcd
4. 如果整体 gcd > 1，返回 -1
5. 否则，寻找最短子数组使 gcd = 1
6. 计算总操作次数：`n + minLen - 2`

## 代码实现

```go
func minOperations(nums []int) int {
    n := len(nums)
    
    // 统计1的个数
    countOnes := 0
    for _, num := range nums {
        if num == 1 {
            countOnes++
        }
    }
    
    // 情况1：已经有1
    if countOnes > 0 {
        return n - countOnes
    }
    
    // 情况2：没有1，检查整体gcd
    overallGcd := nums[0]
    for i := 1; i < n; i++ {
        overallGcd = gcd(overallGcd, nums[i])
    }
    
    if overallGcd > 1 {
        return -1 // 情况3：不可能
    }
    
    // 情况4：寻找最短子数组
    minLen := n + 1
    for i := 0; i < n; i++ {
        currentGcd := nums[i]
        for j := i + 1; j < n; j++ {
            currentGcd = gcd(currentGcd, nums[j])
            if currentGcd == 1 {
                minLen = min(minLen, j - i + 1)
                break
            }
        }
    }
    
    return n + minLen - 2
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 时间复杂度分析

- 统计 1 的个数：O(n)
- 计算整体 gcd：O(n)
- 寻找最短子数组：O(n²)
- 总体时间复杂度：O(n²)

由于题目中 n ≤ 50，O(n²) 的解法是完全可以接受的。

## 关键思考点

1. **为什么会想到检查整体 gcd？**
   - 如果整体 gcd > 1，说明所有数都有共同的因子，无论如何操作都无法消除这个因子

2. **为什么要找最短子数组？**
   - 找到第一个 1 的代价是子数组长度减 1
   - 越短的子数组意味着越早得到 1，总操作次数越少

3. **GCD 的传播特性**
   - 一旦得到一个 1，它就可以像"病毒"一样传播到整个数组
   - 每次传播只需要 1 次操作

## 总结

这道题的核心是理解 GCD 的性质和传播规律。它不是传统的算法题，更像是一道数学思维题。关键在于：

1. 理解 1 在 GCD 运算中的特殊地位
2. 认识到 GCD 的传播特性
3. 通过找最短子数组来优化操作次数

通过这道题，我们可以看到算法题有时更需要数学思维，而不是死板的算法模板。