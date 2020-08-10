package main

import "fmt"

func restoreIpAddresses(s string) []string {
	return part(s, 4)
}

func part(s string, n int) []string {
	short:= n-1
	long := 3*(n-1)

	var pre int
	var head string
	var result []string
	for i:=0;i<3;i++ {
		if i>=len(s) || (pre == 0 && i>0){
			break
		}
		pre = pre * 10 + int(s[i] - '0')
		if (len(s)-i-1) < short || (len(s)-i-1) > long {
			continue
		}
		head = s[:i+1]
		if pre <= 255 {
			if n == 1 {
				result = append(result, head)
			} else {
				left := part(s[i+1:],n-1)
				if left != nil {
					for _, v := range left {
						result = append(result, head+"."+v)
					}
				}
			}
		}
	}
	return result
}

func main()  {
	fmt.Println(restoreIpAddresses("25525511135")) //[255.255.11.135 255.255.111.35]
	fmt.Println(restoreIpAddresses("010010")) //[0.10.0.10 0.100.1.0]
}
