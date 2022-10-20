package main

func backspaceCompare(s string, t string) bool {
	return trim(s) == trim(t)
}

func trim(s string) string {
	var data []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(data) > 0 {
				data = data[:len(data)-1]
			}
		} else {
			data = append(data, s[i])
		}
	}
	return string(data)
}
