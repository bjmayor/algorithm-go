package main

import "fmt"

func solveEquation(equation string) string {
	tmp := 0
	times := 0
	op := '+'
	ratio := 1
	left := true
	v := 0
	for i:=0;i<len(equation);i++ {
		ch := equation[i]
		switch ch {
		case '=':
			if tmp > 0 {
				v = operate(op, v, tmp)
				tmp=0
			}
			left = false
			op = '-'
		case '-':
			if tmp > 0 {
				v = operate(op, v, tmp)
				tmp=0
			}
			if left {
				op = '-'
			} else {
				op = '+'
			}
		case '+':
			if tmp > 0 {
				v = operate(op, v, tmp)
				tmp=0
			}
			if left {
				op = '+'
			} else {
				op = '-'
			}
		case 'x':
			if ratio > 0 {
				if op!='-' {
					times+=ratio
				} else {
					times-=ratio
				}
			}
			tmp=0
			ratio = 1
		default:
			for j:=i;j<len(equation);j++ {
				if equation[j] >='0' && equation[j] <='9' {
					tmp = tmp * 10 + int(equation[j]-'0')
					i=j
				} else {
					if equation[j] == 'x' {
						ratio = tmp
					}
					i = j-1
					break
				}
			}
		}
	}
	if tmp > 0 {
		v = operate(op, v, tmp)
		tmp=0
	}
	if times==0 {
		if v ==0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	return fmt.Sprintf("x=%d",-v/times)
}

func operate(op int32, x, y int) int  {
	switch op {
	case '-':
		return x-y
	case '+':
		return x+y
	}
	return x
}

func main()  {
	//println(solveEquation("x+5-3+x=6+x-2")) //x=2
	//println(solveEquation("x=x")) // Infinite solutions
	//println(solveEquation("2x=x")) // x=0
	//println(solveEquation("1+1=x")) // x=2
	//println(solveEquation("0x=0")) // Infinite solutions
	println(solveEquation("3x=33+22+11")) // x=22
}