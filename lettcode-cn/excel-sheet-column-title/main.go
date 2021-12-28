package main

import (
	"bytes"
)

func convertToTitle(columnNumber int) string {
	var buf bytes.Buffer
	for columnNumber > 0 {
		columnNumber--
		buf.WriteByte(byte(columnNumber%26) + 'A')
		columnNumber /= 26
	}
	temp := buf.Bytes()
	for i:=0;i<len(temp)/2;i++ {
		temp[i],temp[len(temp)-1-i] = temp[len(temp)-1-i],temp[i]
	}
	return string(temp)
}

func main()  {
	println(convertToTitle(1)) // A
	println(convertToTitle(26)) // Z
	println(convertToTitle(52)) // AZ
	println(convertToTitle(27)) // AA
	println(convertToTitle(28)) // AB
	println(convertToTitle(701)) // ZY
	println(convertToTitle(2147483647)) // FXSHRXW
}