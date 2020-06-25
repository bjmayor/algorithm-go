package main

import (
	"fmt"
)

func patternMatching(pattern string, value string) bool {
	if pattern == "" {
		if value == "" {
			return true
		}
		return false
	} else {
		m := map[uint8]int{'a':0,'b':0}
		for i:=0;i<len(pattern);i++ {
			m[pattern[i]]++
		}
		if m['a'] * m['b'] == 0 {
			if value == "" {
				return true
			} else {
				cnt := m['a']+m['b']
				if len(value) % cnt != 0 {
					return false
				}
				size := len(value)/cnt
				substr := value[:size]
				for i:=size;i<len(value);i+=size{
					if value[i:size+i] != substr {
						return false
					}
				}
				return true
			}
		} else {
			if value == "" {
				return false
			}
			//a,b 都有, 3种可能，a为空,b不为空， a不为空，b为空， a,b都不为空
			alla := ""
			for i:=0;i<m['a'];i++ {
				alla += "a"
			}
			if patternMatching(alla, value) {
				return true
			}

			allb := ""
			for i:=0;i<m['b'];i++ {
				allb += "b"
			}
			if patternMatching(allb, value) {
				return true
			}
			first:=m[pattern[0]]
			other := m['a']+m['b']-first
			//a的长度从1到  (len(values)-b的个数)/a的个数
			maxSize := (len(value)-other)/first
			for i:=1;i<=maxSize;i++ {
				if (len(value)-i*first)	 % other !=0 {
					continue
				}
				otherSize := (len(value)-i*first) / other
				fsub := value[0:i]
				start := 0
				for j:=0;j<len(pattern);j++{
					if pattern[j] == pattern[0] {
						start+= i
					}	else {
						break
					}
				}
				osub := value[start:start+otherSize]
				if fsub == osub {
					continue
				}
				start =0
				ok := true
				for j:=0;j<len(pattern);j++{
					if pattern[j] == pattern[0] {
						if value[start:start+i] != fsub {
							ok = false
							break
						}
						start+=i
					}	else {
						if value[start:start+otherSize] != osub{
							ok = false
							break
						}
						start+=otherSize
					}
				}
				if ok {
					return true
				}
			}
			return false

		}
	}
}


func main()  {
	tests := []struct{
		Pattern string
		Value string
		Expected bool
	} {
		//{ "abba","dogcatcatdog",true},
		//{ "abba","dogcatcatfish",false},
		//{ "aaaa","dogcatcatdog",false},
		//{ "abba","dogdogdogdog",true},
		//{ "ab","",false},
		{ "baabab","eimmieimmieeimmiee",true},
	}
	for i, t := range tests {
		var real = patternMatching(t.Pattern, t.Value)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Pattern)
		}
	}
}