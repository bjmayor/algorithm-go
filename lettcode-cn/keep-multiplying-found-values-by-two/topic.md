2154. 将找到的值乘以 2

## 题目描述

给你一个整数数组 `nums` ，另给你一个整数 `original` ，这是需要在 `nums` 中搜索的第一个数字。

接下来，你需要按下述步骤操作：

1. 如果在 `nums` 中找到 `original` ，将 `original` **乘以** 2 ，得到新 `original`（即，令 `original = 2 * original`）。
2. 否则，停止这一过程。
3. 只要能在数组中找到新 `original` ，就对新 `original` 继续 **重复** 这一过程**。**

返回 `original` 的 **最终** 值。

## 示例

**示例 1：**
```
输入：nums = [5,3,6,1,12], original = 3
输出：24
解释：
- 3 能在 nums 中找到。3 * 2 = 6 。
- 6 能在 nums 中找到。6 * 2 = 12 。
- 12 能在 nums 中找到。12 * 2 = 24 。
- 24 不能在 nums 中找到。因此，返回 24 。
```

**示例 2：**
```
输入：nums = [2,7,9], original = 4
输出：4
解释：
- 4 不能在 nums 中找到。因此，返回 4 。
```

## 提示

* `1 <= nums.length <= 1000`
* `1 <= nums[i], original <= 1000`

## 解题思路

### 提示 1
Repeatedly iterate through the array and check if the current value of original is in the array.

### 提示 2
If original is not found, stop and return its current value.

### 提示 3
Otherwise, multiply original by 2 and repeat the process.

### 提示 4
Use set data structure to check the existence faster.

## 相关标签

[数组](/tag/array/) [哈希表](/tag/hash-table/) [排序](/tag/sorting/) [模拟](/tag/simulation/) [第 278 场周赛](/contest/weekly-contest-278/)

## 相似题目

[至少是其他数字两倍的最大数](/problems/largest-number-at-least-twice-of-others/)
简单

[检查整数及其两倍数是否存在](/problems/check-if-n-and-its-double-exist/)
简单
