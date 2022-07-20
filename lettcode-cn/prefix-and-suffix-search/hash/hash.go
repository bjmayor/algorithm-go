package main


type WordFilter struct {
	m map[string]int
}


func Constructor(words []string) WordFilter {
	m := make(map[string]int)
	for idx, word := range words {
		for j:=1;j<=len(word);j++ {
			for k:=0;k<len(word);k++ {
				m[word[:j]+"|"+word[k:]] = idx
			}
		}
	}
	return WordFilter{m:m}
}


func (this *WordFilter) F(pref string, suff string) int {
	if v, ok := this.m[pref+"|"+suff];ok {
		return v
	} else {
		return -1
	}
}




func main()  {
	//obj := Constructor([]string{"apple"});
	//println(obj.F("a","e"));
	obj := Constructor([]string{"abbba","abba"});
	println(obj.F("ab","ba"));
}
