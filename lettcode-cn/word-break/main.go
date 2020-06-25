package main

import (
	"fmt"
)
var m map[string]bool
//dfs + 记忆化
func wordBreak(s string, wordDict []string) bool {
	m = make(map[string]bool)
	trie := Constructor()
	for _,word := range wordDict {
		trie.Insert(word)
	}
	for i:=0;i<len(s);i++ {

	}
	return dfs(s,  &trie)
}

func dfs(s string, trie *Trie) bool  {
	if 0 == len(s) {
		return true
	}
	if v, ok := m[s]; ok {
		return v
	}
	if trie.children[s[0]] == nil {
		m[s] = false
		return false
	}
	cur := trie
	for j:=0;j<len(s);j++ {
		if cur.children[s[j]] !=nil {
			if cur.children[s[j]].isWord {
				if dfs(s[j+1:], trie) == true {
					return true
				}
			}
			cur = cur.children[s[j]]
		}else {
			m[s] = false
			break
		}
	}
	m[s] = false
	return m[s]
}

type Trie struct {
	val byte
	isWord bool
	children map[byte]*Trie
	word string
}


/** Initialize your data structure here. */
func Constructor() Trie {
	root := Trie{}
	root.isWord = false
	root.children = make(map[byte]*Trie)
	return root
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	parent := this
	for i:=0;i<len(word);i++ {
		char := word[i]
		if trie, ok := parent.children[char]; ok {
			parent = trie
		} else {
			child := Constructor()
			child.val = char
			child.word = parent.word+fmt.Sprintf("%c", char)
			parent.children[char] = &child
			parent = &child
		}
	}
	parent.isWord = true
}

//动态规划
func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	for i:=0;i<len(s);i++ {
		for j:=i;j>=0;j-- {
			if m[s[j:i+1]] && dp[j] {
				dp[i+1] = true
				break
			}
		}
	}
	return dp[len(s)]
}


func main()  {
	//fmt.Println(wordBreak2("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",[]string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}))//false
	//fmt.Println(wordBreak2("cars",[]string{"car","ca","rs"}))//true
	//fmt.Println(wordBreak2("leetcode",[]string{"leet","code"}))//true
	//fmt.Println(wordBreak2("applepenapple",[]string{"apple","pen"}))//true
	//fmt.Println(wordBreak2("catsandog",[]string{"cats", "dog", "sand", "and", "cat"}))//false
	fmt.Println(wordBreak2("a",[]string{"a"}))//true
}