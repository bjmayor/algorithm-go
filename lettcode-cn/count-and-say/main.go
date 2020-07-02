package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n== 1 {
		return "1"
	}
	pre := "1"
	for i:=2;i<=n;i++ {
		pc := pre[0]
		cnt :=1
		npre := ""
		for j:=1;j<len(pre);j++ {
			if pre[j] == pc {
				cnt ++
			} else {
				npre += strconv.Itoa(cnt)+fmt.Sprintf("%c", pc)
				pc = pre[j]
				cnt = 1
			}
		}
		npre += strconv.Itoa(cnt)+fmt.Sprintf("%c", pc)
		pre = npre
	}
	return pre
}

func main()  {
	tests := []struct{
		Nums int
		Expected string
	} {
		{ 1,"1"},
		{ 4,"1211"},
	}
	for _, t := range tests {
		var real = countAndSay(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}