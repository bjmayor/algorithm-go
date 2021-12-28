package main


// 试试广度优先
func slidingPuzzle(board [][]int) int {

	var start [2][3]int
	for i:=0;i<2;i++ {
		for j:=0;j<3;j++ {
			start[i][j] = board[i][j]
		}
	}
	var target  = [2][3]int{{1,2,3},{4,5,0}}
	if target == start {
		return 0
	}

	// 枚举 status 通过一次旋转得到的数字
	get := func(status [2][3]int) (ret [][2][3]int) {
		for i:=0;i<2;i++ {
			for j:=0;j<3;j++ {
				if status[i][j] == 0 {
					if i-1 >=0 {
						status[i][j], status[i-1][j] =  status[i-1][j],status[i][j]
						ret = append(ret, copy(status))
						status[i][j], status[i-1][j] =  status[i-1][j],status[i][j]
					}
					if i+1 <=1 {
						status[i][j], status[i+1][j] =  status[i+1][j],status[i][j]
						ret = append(ret, copy(status))
						status[i][j], status[i+1][j] =  status[i+1][j],status[i][j]
					}
					if j-1 >=0 {
						status[i][j], status[i][j-1] =  status[i][j-1],status[i][j]
						ret = append(ret, copy(status))
						status[i][j], status[i][j-1] =  status[i][j-1],status[i][j]
					}
					if j+1 <=2 {
						status[i][j], status[i][j+1] =  status[i][j+1],status[i][j]
						ret = append(ret, copy(status))
						status[i][j], status[i][j+1] =  status[i][j+1],status[i][j]
					}
					break
				}
			}
		}
		return
	}

	type pair struct {
		status [2][3]int
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[[2][3]int]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}

func copy(src [2][3]int) (ret [2][3]int) {
	for i:=0;i<2;i++ {
		for j:=0;j<3;j++ {
			ret[i][j] = src[i][j]
		}
	}
	return
}

func main()  {
	println(slidingPuzzle([][]int{{1,2,3},{4,0,5}})) // 1
	println(slidingPuzzle([][]int{{1,2,3},{5,4,0}})) // -1
	println(slidingPuzzle([][]int{{4,1,2},{5,0,3}})) // 5
	println(slidingPuzzle([][]int{{3,2,4},{1,5,0}})) // 14
}