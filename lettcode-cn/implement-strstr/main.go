package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	next := nextJ(needle)
	i:=0
	j:=0
	for i<len(haystack) && j<len(needle) {
		if needle[j] == haystack[i] {
			j++
			i++
		} else {
			nj := next[j]
			if nj == -1 {
				i++
				j = 0
			} else {
				j = nj
			}
		}
	}
	if j == len(needle) {
		return i-j
	}
	return -1
}

func nextJ(needle string) []int{
	var result = make([]int, len(needle))
	result[0] = -1 //表示首字符匹配失败，i需要后移一位
	j:=1
	k:=-1
	for j < len(needle) {
		if k==-1 ||  needle[j-1] == needle[k] { //1. 找不到匹配，所以才0开始匹配 2. 可以加1
			k++
			if needle[j] != needle[k] { //增加修正
				result[j] = k
			} else {
				result[j] = result[k]
			}
			j++
		} else { //找k
			k = result[k]
		}
	}
	return result
}

func main()  {
	tests := []struct{
		Nums string
		Needle string
		Expected int
	} {
		{ "hello","ll",2},
		{ "aaaaa","bba",-1},
		{ "aabaabdc","aabd",3},
	}
	for _, t := range tests {
		var real = strStr(t.Nums, t.Needle)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
