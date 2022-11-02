package main

func magicalString(n int) int {
	if n < 4 {
		return 1
	}

	s := "122"
	i := 2 // 最后1个
	j := 2 // 下一个
	v := 1
	for len(s) < n {
		if s[j] == '2' {
			if s[i] == '1' {
				s += "22"
			} else {
				s += "11"
				v += 2
				if len(s) == n+1 {
					v -= 1
				}
			}
			i += 2

		} else {
			if s[i] == '1' {
				s += "2"
			} else {
				s += "1"
				v += 1
			}
			i += 1
		}
		j++
	}
	return v
}

func main() {
	println(magicalString(7))
}
