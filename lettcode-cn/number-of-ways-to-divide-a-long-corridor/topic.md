2147. 分隔长廊的方案数
困难
相关标签
premium lock icon
相关企业
提示
在一个图书馆的长廊里，有一些座位和装饰植物排成一列。给你一个下标从 0 开始，长度为 n 的字符串 corridor ，它包含字母 'S' 和 'P' ，其中每个 'S' 表示一个座位，每个 'P' 表示一株植物。

在下标 0 的左边和下标 n - 1 的右边 已经 分别各放了一个屏风。你还需要额外放置一些屏风。每一个位置 i - 1 和 i 之间（1 <= i <= n - 1），至多能放一个屏风。

请你将走廊用屏风划分为若干段，且每一段内都 恰好有两个座位 ，而每一段内植物的数目没有要求。可能有多种划分方案，如果两个方案中有任何一个屏风的位置不同，那么它们被视为 不同 方案。

请你返回划分走廊的方案数。由于答案可能很大，请你返回它对 109 + 7 取余 的结果。如果没有任何方案，请返回 0 。


提示 1
Divide the corridor into segments. Each segment has two seats, starts precisely with one seat, and ends precisely with the other seat.
提示 2
How many dividers can you install between two adjacent segments? You must install precisely one. Otherwise, you would have created a section with not exactly two seats.
提示 3
If there are k plants between two adjacent segments, there are k + 1 positions (ways) you could install the divider you must install.
提示 4
The problem now becomes: Find the product of all possible positions between every two adjacent segments.

