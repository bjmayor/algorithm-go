package main

import "fmt"



func letterCombinations(digits string) []string {
	var m = []string{ "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz", }
	var result []string
	if len(digits) == 0 {
		return nil
	}
	result = append(result,"")

	for i := 0; i < len(digits); i++ {
		size:= len(result)
		for k:=0;k<size;k++ {
			s := result[0]
			result = result[1:]
			t := m[digits[i]-'2']
			for j:=0;j<len(t);j++ {
				result = append(result, s+t[j:j+1])
			}
		}

	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
