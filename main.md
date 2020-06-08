# 算法
自娱自乐的。把见过的一些算法记录下来。

# Fisher-Yates算法
用于把数组打乱。

# 最小可用ID
smallest-avaiable-id,
找出不走某列表的最小的非负整数。
某列表大小可能上千万。

由于列表是不重复的非负整数列表。
那么就有，如果a0, a1,a2,....an 存在某个数 < n 未使用。那么 必然存在某个数ai > n。
否则a0,a1,a2,....an 就是刚好0~n都用了。最小可用id 就是 n+1。
 minfree(a0,a1,....an) <= n+1
 
 由词可得分治法。
 选定一个数ai，切割数组a0~an, 小于等于的放左边，大于的放右边。
 如果小于等于的个数 刚好等于i, 说明左边的满了，再从右边的子集找，否则从左边的子集找。
 
 # 寻找丑数
 ugly-number
 寻找第1500个"丑数"。
 定义：
 丑数 是指只含2、3、5 这3个素因子的自然数。
 
 前十个丑数是：
 2，3，4，5，6，8，9，10，12，15
 
 这里是利用队列实现复杂度为O(n)的算法。

# 并集查
见 lettcode-cn/reverse-integer

# 2点间所有最短路径问题
见 lettcode-cn/word-ladder-li

# 动态规划
lettcode-cn/longest-palindromic-substring

# Manacher算法
专门用于查找最长回文子串的算法。
lettcode-cn/longest-palindromic-substring

# 堆排序
lettcode-cn/kth-largest-element-in-a-stream
703