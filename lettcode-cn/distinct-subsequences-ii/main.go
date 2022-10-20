package main

func distinctSubseqII(s string) (ans int) {
	var dp [26]int
	for _, v := range s {
		sum := 0
		for j := 0; j < 26; j++ {
			sum = (sum + dp[j]) % (1e9 + 7)
		}
		dp[v-'a'] = sum + 1 // 这里就是动态转移方程。。不选是sum, 选是1
	}

	for _, v := range dp {
		ans += v
	}
	return ans % (1e9 + 7)
}

func main() {
	println(distinctSubseqII("abc")) // 7
	println(distinctSubseqII("aaa")) // 3
}
