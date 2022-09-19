package main

func maxLengthBetweenEqualCharacters(s string) int {
	m := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)
	}
	var ans = -1
	for _, v := range m {
		if len(v) > 1 {
			dis := v[len(v)-1] - v[0] - 1
			if dis > ans {
				ans = dis
			}
		}
	}
	return ans
}
