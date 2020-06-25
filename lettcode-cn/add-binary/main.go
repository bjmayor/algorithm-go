package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	var buffer []byte
	var left uint8 = 0
	i:=len(a)-1
	j:=len(b)-1
	k := max(i,j)+1
	buffer = make([]byte,k+1)
	for i>=0 || j >=0 {
		val := left
		if i>=0 {
			val += a[i]-'0'
			i--
		}
		if j>=0 {
			val += b[j]-'0'
			j--
		}
		left = val/2
		buffer[k] = val%2+'0'
		k--
	}
	if left == 1 {
		buffer[0] = '1'
	} else {
		buffer = buffer[1:]
	}
	return string(buffer)
}

func max(x,y int) int  {
	 if x > y {
	 	return x
	 }
	 return y
}

func main()  {
	fmt.Println(addBinary("11","1")) // 100
	fmt.Println(addBinary("1010","1011")) // 10101
	fmt.Println(addBinary("1","111")) // 1000
	fmt.Println(addBinary("100","110010")) // 110110
}