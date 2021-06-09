https://leetcode-cn.com/problems/profitable-schemes/

暴力法 是 C(n,0)+C(n,1)+...C(n,n) = 2的n次方。
(a+b)的n次方的二项式展开式。另a=b=1 即可得出。
时间复杂度太高。

考虑动态规划。
定义dp[i][j][k]为j个员工参与前i个项目的并能产生最少k利润的组合数。
转义方程为：
dp[i+1][j][k] = dp[i][j][k]+dp[i][j-group[i+1]][k-profite[i+1]]
说明dp[i][j][k]表示第i+1个项目不参与。
dp[i][j-group[i+1]][k-profite[i+1]] 表示第i+1个项目参与。

最终的结果就是dp[len(group)][0...n][minProfit] 的和。


时间复杂度是 n*len(group)*minProfit 远小于 2的n次方。

提交可解：
执行用时： 56 ms , 在所有 Go 提交中击败了 61.11% 的用户
内存消耗： 20.8 MB , 在所有 Go 提交中击败了 16.67% 的用户

参考了最佳答案，思路一样，不过是用了滚动数组。
