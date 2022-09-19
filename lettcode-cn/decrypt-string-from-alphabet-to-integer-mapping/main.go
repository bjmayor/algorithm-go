package main

import (
	"bytes"
	"strconv"
)

func freqAlphabets(s string) string {

	var buf bytes.Buffer
	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' {
			v, _ := strconv.Atoi(s[i : i+2])
			buf.WriteByte(byte(v - 10 + 'j'))
			i += 3
		} else {
			v, _ := strconv.Atoi(s[i : i+1])
			buf.WriteByte(byte(v - 1 + 'a'))
			i += 1
		}
	}
	return buf.String()
}
