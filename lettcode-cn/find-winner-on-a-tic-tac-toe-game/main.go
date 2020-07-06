package main
func tictactoe(moves [][]int) string {
	desk := make([][]int, 3, 3)
	for i:=0;i<3;i++ {
		desk[i] = make([]int,3,3)
	}
	for i:=0;i<len(moves);i++ {
		if i & 1 == 1 {
			desk[moves[i][0]][moves[i][1]] = -1
		} else {

			desk[moves[i][0]][moves[i][1]] = 1
		}
	}
	var cnt [8]int
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			cnt[i] += desk[i][j] //行的结果
			cnt[j+3] += desk[i][j] //列的结果
			if i == j {
				cnt[6] += desk[i][j]
			}
			if i+j == 2 {
				cnt[7] += desk[i][j]
			}
		}
	}
	for i:=0;i<8;i++ {
		if cnt[i] == 3 {
			return "A"
		}
		if cnt[i] == -3 {
			return "B"
		}
	}
	if len(moves) == 9 {
		return "Draw"
	} else {
		return "Pending"
	}

}
