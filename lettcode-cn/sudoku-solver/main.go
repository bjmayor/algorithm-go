package main

import "fmt"

func solveSudoku(board [][]byte)  {
	tryFill(0, board)
}

func tryFill(step int, board [][]byte) bool {
	if step == 81 {
		return true
	}
	i, j := step/9, step%9
	if board[i][j] != '.' {
		return tryFill(step+1, board)
	} else {
		for k:=1;k<=9;k++ {
			if isValidSudoku(board,i,j, uint8(k) + '0') {
				board[i][j] = uint8(k) + '0'
				if tryFill(step+1, board) {
					return true
				}
				board[i][j] = '.' //这里是回退
			}
		}
		return false
	}
}


func isValidSudoku(board [][]byte, i,j int, val byte) bool {
	for c:=0;c<9;c++ {
		if board[i][c] == val {
			return false
		}
	}
	for r:=0;r<9;r++ {
		if board[r][j] == val {
			return false
		}
	}

	 r :=  i/3*3
	 c := j/3*3
	 for k:=0;k<3;k++ {
	 	for l:=0;l<3;l++ {
	 		if  board[k+r][l+c] == val {
	 			return false
			}
		}
	 }
	 return true
}

func main()  {

	board2 := [][]byte{
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '6', '.', '.'},
		[]byte{'.', '.', '.', '3', '8', '.', '.', '.', '.'},
		[]byte{'.', '5', '.', '.', '.', '6', '.', '.', '1'},
		[]byte{'8', '.', '.', '.', '.', '.', '.', '6', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '7', '.', '9', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '6', '.', '.', '.', '.', '.'},
	}
	//fmt.Println(isValidSudoku(board))
	solveSudoku(board2)
	fmt.Println(board2)
}