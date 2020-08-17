package main

import "fmt"

func solve(board [][]byte)  {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	for j:=0;j<n;j++ {
		if board[0][j] == 'O' {
			dfs(board, 1, j)
		}
		if board[m-1][j] == 'O' {
			dfs(board, m-2, j)
		}
	}

	for i:=1;i<m-1;i++ {
		if board[i][0] == 'O' {
			dfs(board, i, 1)
		}
		if board[i][n-1] == 'O' {
			dfs(board, i, n-2)
		}
	}

	for i:=1;i<m-1;i++ {
		for j:=1;j<n-1;j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else  {
				board[i][j]= 'X'
			}
		}
	}
}

func dfs(board [][]byte,  x,y  int) {
	if  x>len(board)-2 || x < 1 || y < 1 || y>len(board[x])-2 || board[x][y] != 'O'  {
		return
	}
	board[x][y] = 'A'
	dfs(board,x+1,y)
	dfs(board,x-1,y)
	dfs(board,x,y+1)
	dfs(board,x,y-1)
}

func main()  {
	var board = [][]byte{
		{'O'},

	}
	solve(board)
	fmt.Println(board)
}
