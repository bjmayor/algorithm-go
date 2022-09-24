package main

import "bytes"

func isIsomorphic(s string, t string) bool {
	m := map[byte]byte{}
	n := map[byte]byte{}
	m[t[0]] = s[0]
	n[s[0]] = t[0]
	var buf bytes.Buffer
	buf.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if v, ok := m[t[i]]; ok {
			buf.WriteByte(v)
		} else {
			if _, ok := n[s[i]]; ok {
				return false
			}
			m[t[i]] = s[i]
			n[s[i]] = m[t[i]]
			buf.WriteByte(s[i])
		}
	}
	return s == buf.String()
}

func main() {
	println(isIsomorphic("foo", "bar"))
}
