https://leetcode-cn.com/problems/last-stone-weight-ii/

最后一块石头的重量 II
第一反应是动态规范。稍一琢磨又不像。因为在已经dp[n]的解的情况下，没法获取dp[n+1]的解。

没啥思路的情况下，只能按步骤走一下
输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值

换种顺序看看，就按顺序来：
[2,7,4,1,8,1]
2,7组合 [5,4,1,8,1]
5,4 组合 [1,1,8,1]
1，1 组合[8,1]
8,1 组合[7] 不是最优解。

在已知最优解的情况下，能不能推断出一个组合顺序。
[2,7,4,1,8,1] 已知解为1,
说明[2,7,4,1,8] 的解为0
说明[7,4,1,8]的解为2， 8-7=1, 4-1=3, 3-1=2, 确实有解。
说明[4,1, 8] 的解要么为5，要么为9. 4-1=3, 8-3=5 还真有解。怎么组合呢？？
在只有1个或者2个时候的情况下。他们的解是确定的。
如[8]的解只能是8
[8,4]的解一定是4.
三个呢。
[8,4,2] 分成2堆有3种方法。
1. [8,4],[2] 解为2
2. [8,2], [4] 解为2
3. [2,4], [8] 解为6
划分成2对的规模是2，1 都是确定的解，取最优解即可。
   
4个呢？
[8,4,2,9] 分成2堆有4种方法。 这里就不同了3个一堆的有多解决。
1. [8,4,2],[9], [8,4,2]有2种解2和6， 最终还是2种解:7、3
2. [8,4,9],[2],[8,4,9]的解是:5,3, 最终解是3,1
3. [8,9,2],[4], [8,9,2]的解是1，3， 最终解是3,1
4. [9,4,2],[8], [9,4,2]的解是3, 最终解是5
5. [8,4],[2,9] 解为3
6. [8,2],[4,9] 解为1
7. [8,9],[2,4] 解为1
显然最优解是1, 感觉也只能遍历求解。实现倒也容易。
   就是从里面挑2个石头。之后放回去。
   选2个时候是nx(n-1), 之后问题规模-1 或者 -2。最坏的情况问题规模-1
   f(n) = n*(n-1) * f(n-1)
   =n*(n-1) *(n-1)*(n-2)*f(n-2)
   = ...
   n!的平方。。肯定超时啊。
   
把解法展开。
[8,4,2] 分成2堆有3种方法。
1. [8,4],[2] 解为2 = 8-4-2
2. [8,2], [4] 解为2 8-2-4
3. [2,4], [8] 解为6 8-(4-2) = 8-4+2
可以发现每种解法其实都是在数字前添加了+号或者-号。 这恰好和之前的一道题是一样的。
   https://leetcode-cn.com/problems/target-sum/
   而这道目标和的解法是 转换成多少种方式使得前i个数字之和为target.(target = (target+sum)/2)
   
这道题是要求有不同的组合的和最小。
那就是目标和问题。使target=0->100
如果组合的种数>0, 即有解。就返回。 时间复杂度就是o(nxsum)了。

   

   
   