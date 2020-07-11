package main

import (
	"fmt"
)

//trie树 + dp
func respace(dictionary []string, sentence string) int {
	rtrie := Trie{}
	for _, v := range dictionary {
		rtrie.Insert(v)
	}
	dp := make([]int, len(sentence)+1) //dp[i], 前i字符的解
	for i := 0; i < len(sentence); i++ {
		dp[i+1] = i+1
	}
	for i := 1; i <= len(sentence); i++ {
			dp[i] = dp[i-1] + 1
			root := &rtrie
			c := sentence[i-1]-'a'
			start := i-1
			for root.children[c] != nil {
				root = root.children[c]
				if root.isWord {
					if dp[start] < dp[i] {
						dp[i] = dp[start]
					}
				}
				start--
				if start >= 0 {
					c = sentence[start]-'a'
				} else {
					break
				}
			}
	}
	return dp[len(sentence)]
}


type Trie struct {
	isWord   bool
	children [26]*Trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for i := len(word)-1; i >=0;  i-- {
		idx := word[i]-'a'
		if next := parent.children[idx]; next == nil {
			parent.children[idx] = &Trie{}
		}
		parent = parent.children[idx]
	}
	parent.isWord = true
}

func main() {
	fmt.Println(respace([]string{"qqqqqqqqq","qqqqq","qq","qqq","qqqqqq","qqqqqq","q"}, "qqqq"))       //0
	fmt.Println(respace([]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother"))                                                                                                                                                                                                                                                                                                                                                                                                                                //7
	fmt.Println(respace([]string{"sssjjs", "hschjf", "hhh", "fhjchfcfshhfjhs", "sfh", "jsf", "cjschjfscscscsfjcjfcfcfh", "hccccjjfchcffjjshccsjscsc", "chcfjcsshjj", "jh", "h", "f", "s", "jcshs", "jfjssjhsscfc"}, "sssjjssfshscfjjshsjjsjchffffs"))                                                                                                                                                                                                                                                                                            //7
	fmt.Println(respace([]string{"aaysaayayaasyya", "yyas", "yayysaaayasasssy", "yaasassssssayaassyaayaayaasssasysssaaayysaaasaysyaasaaaaaasayaayayysasaaaa", "aya", "sya", "ysasasy", "syaaaa", "aaaas", "ysa", "a", "aasyaaassyaayaayaasyayaa", "ssaayayyssyaayyysyayaasaaa", "aya", "aaasaay", "aaaa", "ayyyayssaasasysaasaaayassasysaaayaassyysyaysaayyasayaaysyyaasasasaayyasasyaaaasysasy", "aaasa", "ysayssyasyyaaasyaaaayaaaaaaaaassaaa", "aasayaaaayssayyaayaaaaayaaays", "s"}, "asasayaayaassayyayyyyssyaassasaysaaysaayaaaaysyaaaa")) //7
}
