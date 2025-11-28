package main

func maxOperations(s string) int {
	result := 0
	ones := 0
	
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ones++
		} else if i > 0 && s[i-1] == '1' {
			// 当前是0，前一个是1，说明找到了一段可以让左边所有1操作一次的位置
			result += ones
		}
	}
	
	return result
}
