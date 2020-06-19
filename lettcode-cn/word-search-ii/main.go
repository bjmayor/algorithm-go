package main

import "fmt"

var result map[string]bool
func findWords(board [][]byte, words []string) []string {
	result = make(map[string]bool,0)
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	visited := make([][]bool, len(board), len(board))
	for i:=0;i<len(board);i++ {
		visited[i] = make([]bool, len(board[i]), len(board[i]))
	}
	for _, child := range trie.children {
		for i:=0;i<len(board);i++ {
			for j:=0;j<len(board[i]);j++ {
				if board[i][j] == child.val {
					dfs(board, visited, i, j, child)
				}
			}
		}

	}
	t := []string{}
	for k, _ := range result {
		t = append(t, k)
	}
	return t
}

func dfs(board [][]byte, visited [][]bool, i,j int, trie *Trie) {
	if !visited[i][j] && board[i][j] == trie.val {
		visited[i][j] = true
		if trie.isWord {
			result[trie.word] = true
			if len(trie.children) == 0 {
				visited[i][j] = false
				return
			}
		}
		for _, child := range trie.children {
			//相邻的4个
			if i-1>=0 && !visited[i-1][j] {
				  dfs(board, visited, i-1, j, child)
			}

			if i+1 < len(board) && !visited[i+1][j] {
				  dfs(board, visited, i+1, j, child)
			}

			if j-1>=0 && !visited[i][j-1] {
				 dfs(board, visited, i, j-1, child)
			}

			if j+1 < len(board[i]) && !visited[i][j+1] {
				  dfs(board, visited, i, j+1, child)
			}
		}
	}
	visited[i][j] = false
	return
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

func main()  {
	//board := [][]byte {
	//		{'o','a','a','n'},
	//		{'e','t','a','e'},
	//		{'i','h','k','r'},
	//		{'i','f','l','v'},
	//}
	//fmt.Println(findWords(board, []string{"oath","pea","eat","rain"})) //true
	//
	//board := [][]byte {
	//	{'a','a'},
	//}
	//fmt.Println(findWords(board, []string{"aa"})) //true


	board := [][]byte {
		{'a','b'},
		{'a','a'},
	}
	fmt.Println(findWords(board, []string{"aba","baa","bab","aaab","aaa","aaaa","aaba"})) //true
}
