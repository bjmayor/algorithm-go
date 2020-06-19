package main

import "fmt"

type Trie struct {
	val byte
	isWord bool
	children map[byte]*Trie

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
			parent.children[char] = &child
			parent = &child
		}
	}
	parent.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for i:=0;i<len(word);i++ {
		char := word[i]
		if trie, ok := parent.children[char]; ok {
			parent = trie
		} else {
			return false
		}
	}
	return parent.isWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	parent := this
	for i:=0;i<len(prefix);i++ {
		char := prefix[i]
		if trie, ok := parent.children[char]; ok {
			parent = trie
		} else {
			return false
		}
	}
	return true
}

func main()  {
	trie := Constructor();
	trie.Insert("apple");
	fmt.Println(trie.Search("apple"))   // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")); // 返回 true
	trie.Insert("app");
	fmt.Println(trie.Search("app"));     // 返回 true

}

