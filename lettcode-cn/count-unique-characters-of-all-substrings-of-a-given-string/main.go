package main

func uniqueLetterString(s string) int {
	var ans int
	m := map[rune][]int{}
	for i, ch := range s {
		m[ch] = append(m[ch], i)
	}
	for _, arr := range m {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return ans
}

func main() {
	println(uniqueLetterString("AAA"))      // 3
	println(uniqueLetterString("ABC"))      // 10
	println(uniqueLetterString("ABA"))      // 8
	println(uniqueLetterString("LEETCODE")) // 92
}
