package main

import "fmt"

func divisorGame(N int) bool {
	return N%2 == 0
}



func main()  {
	tests := []struct{
		Nums int
		Expected bool
	} {
		//{ 2,true},
		//{ 3,false},
		{ 5,false},
	}
	for i, t := range tests {
		var real = divisorGame(t.Nums)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
