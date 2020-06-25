package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	var row, col, dir int
	var result []byte

	data := make([][]byte, numRows, numRows)
	for i:=0;i<numRows;i++ {
		data[i]= make([]byte, 0)
	}
	data[0] = append(data[0], s[0])
	for i:=1;i<len(s);i++ {
		if dir == 0 { //向下
			row++
		} else { //向上
			row--
			col++
		}
		if row  ==  numRows -  1 {
			//到底了, 向上
			dir = 1
		} else if row == 0 {
			dir = 0
		}
		data[row] = append(data[row], s[i])
	}
	for i:=0;i<numRows;i++ {
		result = append(result, data[i]...)
	}
	return string(result)
}

func main()  {
	tests := []struct{
		s string
		numRows int
		Expected string
	} {
		//{ "LEETCODEISHIRING",3,"LCIRETOESIIGEDHN"},
		//{ "LEETCODEISHIRING",4,"LDREOEIIECIHNTSG"},
		{ "A",1,"A"},
		{ "ABCD",2,"ACBD"},
	}
	for i, t := range tests {
		var real = convert(t.s, t.numRows)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real)
		}
	}
}