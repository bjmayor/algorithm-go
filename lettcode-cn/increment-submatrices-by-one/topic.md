2536. 子矩阵元素加 1

给你一个正整数 n ，表示最初有一个 n x n 的矩阵，每个元素都是 0 。

给你一个二维整数数组 queries ，其中 queries[i] = [row1i, col1i, row2i, col2i] 。对于每个查询 i ，请你找到所有满足 row1i <= x <= row2i 且 col1i <= y <= col2i 的位置 (x, y) ，并将这个位置上的元素加 1 。

请你返回执行完所有查询后得到的矩阵。

约束条件：

1 <= n <= 500

1 <= queries.length <= 10^4

0 <= row1i <= row2i < n

0 <= col1i <= col2i < n
