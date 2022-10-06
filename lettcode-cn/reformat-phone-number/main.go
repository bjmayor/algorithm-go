package main

import (
	"bytes"
	"strings"
)

func reformatNumber(number string) string {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	var buf bytes.Buffer
	for n := len(number); n > 4; {
		buf.WriteString(number[:3])
		buf.WriteByte('-')
		number = number[3:]
		n = len(number)
	}
	if len(number) == 4 {
		buf.WriteString(number[:2])
		number = number[2:]
		buf.WriteByte('-')
		buf.WriteString(number)
	} else {
		buf.WriteString(number)
	}
	return buf.String()
}

func main() {
	println(reformatNumber("1-23-45 6"))
	println(reformatNumber("123 4-567"))
	println(reformatNumber("123 4-5678"))
	println(reformatNumber("12"))
}
