package main

import "fmt"

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board), len(board))
	for i:=0;i<len(board);i++ {
		visited[i] = make([]bool, len(board[i]), len(board[i]))
	}
	for i:=0;i<len(board);i++ {
		for j:=0;j<len(board[i]);j++ {
			if board[i][j] == word[0] {
				if tryMatch(board, visited, i, j, 0, word) {
					return true
				} else {
					visited[i][j] = false
				}
			}
		}
	}
	return false
}

func tryMatch(board [][]byte, visited [][]bool, i,j,k int, word string) bool  {
	if k == len(word) {
		return true
	}
	if !visited[i][j] && board[i][j] == word[k] {
		visited[i][j] = true
		if k+1 == len(word) {
			return true
		}
		//相邻的4个
		if i-1>=0 && !visited[i-1][j] {
			if  tryMatch(board, visited, i-1, j, k+1, word) {
				return true
			}
		}

		if i+1 < len(board) && !visited[i+1][j] {
			if  tryMatch(board, visited, i+1, j, k+1, word) {
				return true
			}
		}

		if j-1>=0 && !visited[i][j-1] {
			if  tryMatch(board, visited, i, j-1, k+1, word) {
				return true
			}
		}

		if j+1 < len(board[i]) && !visited[i][j+1] {
			if  tryMatch(board, visited, i, j+1, k+1, word) {
				return true
			}
		}
		visited[i][j] = false
		return false
	}
	visited[i][j] = false
	return false
}

func main()  {
	//board := [][]byte {
	//		{'A','B','C','E'},
	//		{'S','F','C','S'},
	//		{'A','D','E','E'},
	//		}
	//fmt.Println(exist(board, "ABCCED")) //true
	//fmt.Println(exist(board, "SEE")) //true
	//fmt.Println(exist(board, "ABCB")) //false


	//board := [][]byte {
	//	{'a','b'},
	//}
	//fmt.Println(exist(board, "ab")) //true

	board := [][]byte {
		{'a','a','a','a'},
		{'a','a','a','a'},
		{'a','a','a','a'},
	}
	fmt.Println(exist(board, "aaaaaaaaaaaaa")) //false
}