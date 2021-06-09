package main

var	tests = []struct{
	Strs []string
	M int
	N int
	Expected int
}{
	{[]string{"10", "0001", "111001", "1", "0"}, 5,3,4},
	{[]string{"10", "0", "1"}, 1,1,2},
}
