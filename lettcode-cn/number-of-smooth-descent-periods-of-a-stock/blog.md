 思路（修正）

 一次遍历，维护当前平滑下降段的长度 `size`，以及答案 `ans`。初始化 `size = 1`，`ans = 0`。

 遍历 i 从 1 到 n-1：
 - 如果 `prices[i] == prices[i-1] - 1`，则 `size++`（在同一平滑下降段内）；
 - 否则：把当前段贡献的子段数量累加到 `ans`：`ans += size * (size + 1) / 2`，然后 `size = 1` 重新开始。

 遍历结束后别忘了把最后一段的贡献再加一次：`ans += size * (size + 1) / 2`。

 说明：长度为 `k` 的连续段包含的平滑下降子段数等于 1 + 2 + ... + k = `k*(k+1)/2`。

 复杂度：时间 O(n)，空间 O(1)。

 示例：`[3,2,1,4]`，`[3,2,1]` 长度 3 → 贡献 6，`[4]` 长度 1 → 贡献 1，总计 7。

 Go 示例实现：

 ```go
 package main

 import "fmt"

 func countSmoothDescent(prices []int) int64 {
	 n := len(prices)
	 if n == 0 {
		 return 0
	 }
	 var ans int64 = 0
	 size := 1
	 for i := 1; i < n; i++ {
		 if prices[i] == prices[i-1]-1 {
			 size++
		 } else {
			 ans += int64(size * (size + 1) / 2)
			 size = 1
		 }
	 }
	 ans += int64(size * (size + 1) / 2)
	 return ans
 }

 func main() {
	 fmt.Println(countSmoothDescent([]int{3, 2, 1, 4})) // 7
 }
 ```

