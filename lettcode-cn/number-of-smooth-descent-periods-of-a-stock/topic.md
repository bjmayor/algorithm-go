2110. 股票平滑下跌阶段的数目
中等
相关标签
premium lock icon
相关企业
提示
给你一个整数数组 prices ，表示一支股票的历史每日股价，其中 prices[i] 是这支股票第 i 天的价格。

一个 平滑下降的阶段 定义为：对于 连续一天或者多天 ，每日股价都比 前一日股价恰好少 1 ，这个阶段第一天的股价没有限制。

请你返回 平滑下降阶段 的数目。

 

示例 1：

输入：prices = [3,2,1,4]
输出：7
解释：总共有 7 个平滑下降阶段：
[3], [2], [1], [4], [3,2], [2,1] 和 [3,2,1]
注意，仅一天按照定义也是平滑下降阶段。
示例 2：

输入：prices = [8,6,7,7]
输出：4
解释：总共有 4 个连续平滑下降阶段：[8], [6], [7] 和 [7]
由于 8 - 6 ≠ 1 ，所以 [8,6] 不是平滑下降阶段。
示例 3：

输入：prices = [1]
输出：1
解释：总共有 1 个平滑下降阶段：[1]
 

提示：

1 <= prices.length <= 105
1 <= prices[i] <= 105
 
面试中遇到过这道题?
1/5
是
否
通过次数
17,746/30.6K
通过率
58.0%


提示 1
Any array is a series of adjacent longest possible smooth descent periods. For example, [5,3,2,1,7,6] is [5] + [3,2,1] + [7,6].
提示 2
Think of a 2-pointer approach to traverse the array and find each longest possible period.
提示 3
Suppose you found the longest possible period with a length of k. How many periods are within that period? How can you count them quickly? Think of the formula to calculate the sum of 1, 2, 3, ..., k.