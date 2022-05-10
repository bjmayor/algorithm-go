https://leetcode-cn.com/problems/form-largest-integer-with-digits-that-add-up-to-target/

感觉又是送分题。
和硬币问题，01背包问题类似。一看就是dp求解。

定义:dp[i][j]为 使用前i个元素其代价为j的最大整数。
转移方程为：
dp[i+1][j] = max(dp[i][j](不使用第i+1个元素), dp[i][j-cost[i+1]]*10+i+2)

考虑到答案可能很大。所以定义一个字符串表示即可。

考虑到添加的位数放到前面数字更大。所以遍历的cost时，从后往前遍历。

考虑到字符串匹配太废资源。
其实每次只需要考虑第i位有没有选择即可。把他定义为价值为1. 最后比较的就是 位数。
位数越大的有最后结果。
然后再想法生成字符串即可。