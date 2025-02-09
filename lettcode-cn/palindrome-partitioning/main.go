package main

import "fmt"

/*
* partition 将字符串分割成所有可能的回文子串组合
* @param s 输入字符串
* @return [][]string 所有可能的回文子串组合
*
* 解题思路:
* 1. 使用回溯(DFS)搜索所有可能的分割方案
* 2. 对于每个起始位置,尝试所有可能的分割长度
* 3. 判断分割出的子串是否为回文串
* 4. 如果是回文串,继续搜索剩余部分
*/
func partition(s string) [][]string {
	res := [][]string{}
	dfs(s, 0, []string{}, &res)
	return res
}

func dfs(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s[start : i+1]) {
			path = append(path, s[start:i+1])
			dfs(s, i+1, path, res)
			path = path[:len(path)-1] // 回溯,恢复状态
		}
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
