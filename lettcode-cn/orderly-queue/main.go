package main

import "sort"


func orderlyQueue(s string, k int) string {
	if k>1 {
		t:=[]byte(s)
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		return string(t)
	}
	ans := s
	for i := 1; i < len(s); i++ {
		s = s[1:] + s[:1]
		if s < ans {
			ans = s
		}
	}
	return ans

}


func main()  {
	println(orderlyQueue("cba",1)) // acb
	println(orderlyQueue("baaca",3)) // aaabc
}
