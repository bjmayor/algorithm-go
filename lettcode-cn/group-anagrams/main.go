package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		// 将字符串按字母排序
		sortedStr := sortString(str)
		// 将排序后的字符串作为键，将原始字符串加入到映射的值中
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	// 将映射中的所有值收集到结果中
	result := make([][]string, 0, len(anagrams))
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

// 辅助函数：将字符串中的字符按字母顺序排序
func sortString(s string) string {
	slice := strings.Split(s, "")
	sort.Strings(slice)
	return strings.Join(slice, "")
}

func main() {
	// 示例 1
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result1 := groupAnagrams(strs1)
	fmt.Println(result1)

	// 示例 2
	strs2 := []string{""}
	result2 := groupAnagrams(strs2)
	fmt.Println(result2)

	// 示例 3
	strs3 := []string{"a"}
	result3 := groupAnagrams(strs3)
	fmt.Println(result3)
}
