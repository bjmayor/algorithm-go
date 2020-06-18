package main

import "fmt"

type Queen struct {
	row int
	coloun int
	pre *Queen
}
func totalNQueens(n int) int{
	a := []*Queen{}
	result := 0

	//第0行
	for j:=0;j<n;j++ {
		a = append(a, &Queen{
			row:    0,
			coloun: j,
		})
	}
	for len(a) > 0 {
		preQue := a[len(a)-1]

		a = a[:len(a)-1]
		if preQue.row == n-1 {
			result += 1
		} else {
			for j:=0;j<n;j++ {
				tpreQue := preQue
				canUse := true
				for tpreQue != nil {
					if j == tpreQue.coloun || (tpreQue.row+tpreQue.coloun) == (preQue.row+1+j) || (preQue.row+1-tpreQue.row)==(j-tpreQue.coloun) {
						canUse = false
						break
					} else {
						tpreQue = tpreQue.pre
					}
				}
				if canUse {
					a = append(a, &Queen{
						row:    preQue.row+1,
						coloun: j,
						pre:    preQue,
					})
				}
			}
		}

	}
	return result
}

func readable(q *Queen, n int) []string  {
	result := make([]string, n, n)
	for idx := n-1;idx>=0;idx-- {
		result[idx] = genstr(q.coloun, n)
		q = q.pre
	}
	return result
}

func genstr(column,n int) string  {
	var result string
	for j:=0;j<n;j++ {
		if j != column {
			result += "."
		}	else {
			result += "Q"
		}
	}
	return result
}

func main()  {
	fmt.Println(totalNQueens(4))
}
