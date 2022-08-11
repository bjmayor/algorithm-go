package main

import "unicode"

func reformat(s string) string {
	nums := 0
	for _, ch := range s {
		if unicode.IsDigit(ch){
			nums++
		}
	}
	diff := nums -  (len(s)-nums)
	digitsMore := diff>=0
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		return ""
	}
	var res = []byte(s)
	for i,j:=0,1;i<len(res);i+=2{
		// i 需要是类型多的，j需要不是类型多的
		if unicode.IsDigit(rune(res[i])) != digitsMore { // 如果 i不是类型多的，就需要找个交换
			for unicode.IsDigit(rune(res[j])) != digitsMore { // j 找个是类型多的，以和i交换。
				j += 2
			}
			res[i], res[j] = res[j], res[i]
		}
	}
	return string(res)
}
func main()  {
	println(reformat("3ab4d")) // d3a4b
}
