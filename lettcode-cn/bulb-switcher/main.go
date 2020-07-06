package main

import "fmt"

func bulbSwitch(n int) int {
	if n == 0 {
		return 0
	}
	var res int
	for i:=1;i<=n;i++ {
		//最后留下的一定是某个数的平方，这个数经历1/3次变化，从关到开
		if i*i <=n {
			res += 1
		} else {
			break
		}
	}
	return res
}

func main()  {
	tests := []struct{
		Nums int
		Expected int
	} {
		{ 0,0},
		{ 1,1},
		{ 2,1},
		{ 3,1},
		{ 4,2},
		{ 5,2},
	}
	for _, t := range tests {
		var real = bulbSwitch(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
