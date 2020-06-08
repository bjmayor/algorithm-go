package main

import "fmt"

type Union struct {
	ops map[uint8]*map[uint8]bool
}

func NewUnion() *Union {
	u := Union{}
	u.ops = make(map[uint8]*map[uint8]bool)
	return &u
}

func (u *Union)join(x, y uint8)  {
	if u.ops[x] == u.ops[y]{
		if u.ops[x] == nil {
			m := make(map[uint8]bool)
			m[x] = true
			m[y] = true
			u.ops[x] = &m
			u.ops[y] = &m
		}
	} else {
		if u.ops[x]	== nil {
			(*u.ops[y])[x] = true
			u.ops[x] = u.ops[y]

		} else if u.ops[y]	== nil{
			(*u.ops[x])[y] = true
			u.ops[y] = u.ops[x]
		} else {
			for k,_ := range *u.ops[x] {
				(*u.ops[y])[k] = true
			}
			for k,_ := range *u.ops[y] {
				u.ops[k] = u.ops[y]
			}
		}
	}

}

func (u *Union)sameSet(x, y uint8) bool {
	return u.ops[x] == u.ops[y] && u.ops[x]!=nil
}

func equationsPossible(equations []string) bool {
	u := NewUnion()
	for _, eq := range equations {
		op1 := eq[0]
		op := eq[1]
		op2 := eq[3]
		if op == '=' {
			u.join(op1, op2)
		}
	}

	for _, eq := range equations {
		op1 := eq[0]
		op := eq[1]
		op2 := eq[3]
		if op == '!' {
			if op1 == op2 {
				return false
			}
			if u.sameSet(op1,op2) {
				return false
			}
		}
	}
	return true
}

func main()  {
	tests := []struct{
		Nums []string
		Expected bool
	} {
		{ []string{"a==b","b!=a"} ,false},
		{ []string{"a==b","b==a"} ,true},
		{ []string{"a==b","b==c","a==c"} ,true},
		{ []string{"a==b","b!=c","c==a"} ,false},
		{ []string{"c==c","b==d", "x!=z"} ,true},
		{ []string{"a!=a"} ,false},
		{ []string{"a==b","e==c","b==c","a!=e"} ,false},
		{ []string{"e!=c","b!=b","b!=a","e==d"} ,false},
	}
	for _, t := range tests {
		if equationsPossible(t.Nums) !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", equationsPossible(t.Nums), t.Nums)
		}
	}
}
