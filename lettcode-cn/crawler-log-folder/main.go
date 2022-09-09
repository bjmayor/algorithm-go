package main

func minOperations(logs []string) int {
	s := []string{}
	for _, v := range logs {
		if v[0] == '.' {
			if v[1] == '.' {
				if len(s) > 0 {
					s = s[:len(s)-1]
				}
			}
		} else {
			s = append(s, v)
		}
	}
	return len(s)
}
