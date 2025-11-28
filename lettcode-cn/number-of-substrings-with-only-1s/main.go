package main

func numSub(s string) int {
	const MOD = 1000000007
	result := 0
	count := 0

	for _, char := range s {
		if char == '1' {
			count++
		} else {
			result = (result + count*(count+1)/2) % MOD
			count = 0
		}
	}

	// 处理字符串末尾可能存在的连续 '1'
	result = (result + count*(count+1)/2) % MOD

	return result
}
