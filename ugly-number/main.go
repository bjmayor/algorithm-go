package main

import "fmt"

func uglyNumber(n int) int  {
	if n == 1 {
		return 1
	}
	v2 := []int{2}
	v3 := []int{3}
	v5 := []int{5}
	var num int
	for i:=2;i<=n ;i++  {
		min := v2[0]
		if v3[0] < min {
			min = v3[0]
			if v5[0] < min {
				min = v5[0]
				v5 = v5[1:]
				v5 = append(v5, min*5)
			} else {
				v3 = v3[1:]
				v3 = append(v3, min*3)
				v5 = append(v5, min*5)
			}

		} else if v5[0] < min {
			min = v5[0]
			v5 = v5[1:]
			v5 = append(v5, min*5)
		} else {
			v2 = v2[1:]
			v2 = append(v2, 2*min)
			v3 = append(v3, 3*min)
			v5 = append(v5, 5*min)
		}
		num = min
		//fmt.Println(num)
	}
	return num
}

func main()  {
	tests := []struct{
		Nums int
		Expected int
	} {
		{ 1500 ,859963392 },
	}
	for i, t := range tests {
		var real = uglyNumber(t.Nums)
		if real !=  t.Expected {
			fmt.Println(i, " expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
