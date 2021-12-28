https://leetcode-cn.com/problems/binary-subarrays-with-sum/

有类似的题，
做法就是 遍历
当遍历到j时，
存下0..j 的和 tj
0..i 的和(i<=0<j), 用map存下来。
然后查找tj-ti = gaol 就表示 i-j的和为goal 找到一个子数组。
 tj-goal = ti
直接查map, 存在key为ti的元素，就找到了。。


算法复杂度只有O(n),
空间复杂度也是O(n)。

可解，对应的官方的哈希表方法。

官方还提出了一个滑动窗口方法，时间复杂度也是O(n), 空间复杂度只要O(1)
需要注意： 之所以滑动窗口能解决是有前提的： num[i] 不是0 就是1 
如果有负数，这个方法讲不行了。