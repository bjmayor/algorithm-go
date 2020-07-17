package main

import (
	"fmt"
)

type NumCnt struct {
	char uint8
	times int
}

func quickSort(m []NumCnt, left, right int)   {
	if right-left<=1 {
		return
	}
	k := m[left]
	i:=left+1
	j:=right-1
	for i<=j {
		if m[i].times > k.times {
			m[i],m[j] = m[j],m[i]
			j--
		} else {
			i++
		}
		if i>j {
			break
		}
		if m[j].times < k.times {
			m[i],m[j] = m[j],m[i]
			i++
		} else {
			j--
		}
	}
	m[i-1],m[left] = m[left],m[i-1]
	quickSort(m, left, i-1)
	quickSort(m, i, right)
}

func reorganizeString(S string) string {
	m := make(map[uint8]NumCnt, 26)
	for i:=0;i<len(S);i++{
		cnt := m[S[i]]
		cnt.times++
		m[S[i]] = cnt
	}
	var res []uint8
	var cntArr []NumCnt
	for k, v := range m {
		if v.times > (len(S)+1)/2 {
			return ""
		}
		v.char = k
		cntArr = append(cntArr, v)
	}
	quickSort(cntArr, 0, len(m))
	for len(res) < len(S) {
		top := cntArr[len(cntArr)-1]
		for	i:=0;i<top.times-1;i++ {
			res = append(res, top.char)
			for j:=len(cntArr)-2;j>=0;j--{
				if cntArr[j].times > 0 {
					res = append(res, cntArr[j].char)
					cntArr[j].times--
					break
				}
			}
		}
		res = append(res, top.char)
		j := len(cntArr)-2
		for j>=0 && cntArr[j].times == 0  {
			j--
		}
		cntArr = cntArr[:j+1]
		if j>0 {
			tmp := cntArr[j]

			for k:=j-1;k>=0;k-- {
				if cntArr[k].times > tmp.times {
					cntArr[k+1] = cntArr[k]
				} else {
					cntArr[k+1]= tmp
					break
				}
			}
			if cntArr[0].times > tmp.times {
				cntArr[0] = tmp
			}

		}

	}


	return string(res)
}

func isOk(s string) bool  {
	if len(s) == 0 {
		return true
	}
	pre := s[0]
	for i:=1;i<len(s);i++ {
		if s[i] == pre {
			return false
		}
		pre = s[i]
	}
	return true
}

func main()  {
	tests := []struct{
		Nums string
		Expected string
	} {
		//{ "vvvlo","vlvov"},
		//{ "aab","aba"},
		//{ "aaab",""},
		//{ "bfrbs","brbsf"},
		{ "gpneqthatplqrofqgwwfmhzxjddhyupnluzkkysofgqawjyrwhfgdpkhiqgkpupgdeonipvptkfqluytogoljiaexrnxckeofqojltdjuujcnjdjohqbrzzzznymyrbbcjjmacdqyhpwtcmmlpjbqictcvjgswqyqcjcribfmyajsodsqicwallszoqkxjsoskxxstdeavavnqnrjelsxxlermaxmlgqaaeuvneovumneazaegtlztlxhihpqbajjwjujyorhldxxbdocklrklgvnoubegjrfrscigsemporrjkiyncugkksedfpuiqzbmwdaagqlxivxawccavcrtelscbewrqaxvhknxpyzdzjuhvoizxkcxuxllbkyyygtqdngpffvdvtivnbnlsurzroxyxcevsojbhjhujqxenhlvlgzcsibcxwomfpyevumljanfpjpyhsqxxnaewknpnuhpeffdvtyjqvvyzjeoctivqwann","brbsf"},
	}
	for i, t := range tests {
		var real = reorganizeString(t.Nums)
		if real !=  t.Expected && !isOk(real) {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}