package main

import "fmt"

func ambiguousCoordinates(s string) []string {
	tmp := [][2]string{}
	for i := 2; i < len(s)-1; i++ {
		tmp = append(tmp, [2]string{s[1:i], s[i : len(s)-1]})
	}
	var result []string
	for _, s := range tmp {
		result = append(result, combine(s[0], s[1])...)
	}
	return result
}
func combine(s1, s2 string) []string {
	first := possible(s1)
	if len(first) == 0 {
		return nil
	}
	second := possible(s2)
	if len(second) == 0 {
		return nil
	}
	if len(first) > 0 && len(second) > 0 {
		var result = []string{}
		for _, f := range first {
			for _, s := range second {
				result = append(result, "("+f+", "+s+")")
			}
		}
		return result
	}
	return nil
}

func possible(s string) []string {
	var result []string
	if s == "0" {
		return []string{s}
	}

	if s[0] != '0' {
		result = append(result, s)
		if s[len(s)-1] != '0' {
			for i, _ := range s {
				if i != len(s)-1 {
					result = append(result, s[:i+1]+"."+s[i+1:])
				}
			}
		}

	} else {
		if s[len(s)-1] == '0' {
			return nil
		}
		result = append(result, "0."+s[1:])
	}

	return result
}

func main() {
	fmt.Println(ambiguousCoordinates("(123)"))
	fmt.Println(ambiguousCoordinates("(00011)"))
	fmt.Println(ambiguousCoordinates("(0123)"))
	fmt.Println(ambiguousCoordinates("(100)"))
}
