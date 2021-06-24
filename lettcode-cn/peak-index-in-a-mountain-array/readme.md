https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/

送分题。

简单的思路就是遍历，遇到第一次下降时，就找到了i。 O(n)的复杂度。

注意到任选一个元素j 
只存在3种可能。
j-1<j<j+1  上升
j-1>j>j+1 下降
j-1<j>j+1 j即为目标i
于是可以用二分查找法。 O(lgn) 
