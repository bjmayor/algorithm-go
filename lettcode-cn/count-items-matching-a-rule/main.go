package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var ans int
	match := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	for _, item := range items {
		if item[match[ruleKey]] == ruleValue {
			ans++
		}
	}
	return ans
}
