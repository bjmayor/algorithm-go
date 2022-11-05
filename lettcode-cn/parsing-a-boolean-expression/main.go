package main

func parseBoolExpr(expression string) bool {
	stk := []rune{}
	for _, c := range expression {
		if c == ',' {
			continue
		}
		if c != ')' {
			stk = append(stk, c)
			continue
		}
		t := 0
		f := 0
		for stk[len(stk)-1] != '(' {
			val := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if val == 't' {
				t++
			} else {
				f++
			}
		}
		stk = stk[:len(stk)-1]
		op := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		c = 't'
		switch op {
		case '!':
			if f != 1 {
				c = 'f'
			}
			stk = append(stk, c)
		case '&':
			if f != 0 {
				c = 'f'
			}
			stk = append(stk, c)
		case '|':
			if t == 0 {
				c = 'f'
			}
			stk = append(stk, c)
		}
	}
	return stk[len(stk)-1] == 't'
}

func main() {
	println(parseBoolExpr("!(f)"))
}
