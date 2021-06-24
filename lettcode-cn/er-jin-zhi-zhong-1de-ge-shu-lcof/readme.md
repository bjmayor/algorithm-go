https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/

送分题。
1. 是直接%2 取余 1的个数。然后/2. O(n)
2. 是直接右移操作。比1性能高, O(n)
3. 按位与，有多少个1，就操作多少次。 O(m), m为1的个数

