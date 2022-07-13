https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/

暴力法容易实现。遍历累加。
复杂度有点高。O(m x n x indices.length)

像棋盘一样，统计结束以后，为奇数的行或者列刷白，那么不考虑重合就是 sum(r) * n + sum(c) * m，考虑到重合以及刷两次变黑，所以减去两倍的白行 x 白列
