package main

import "fmt"

func findSmallerst(data []int) int  {
	return findSmallerstBase(data, 0)
}

func findSmallerstBase(data []int, base int) int  {
	if len(data) == 0 {
		return base
	}
	if len(data)==1 {
		if data[0] != base {
			return base
		}
		return base+1
	}
	key := data[0]
	i := 1//i之前的都是<=data[0]的数
	j := len(data)-1//j之后的都是大于data[0]的数

	for i<=j {
		for ;i<=j;i++ {
			if data[i] > key {
				data[i], data[j] = data[j], data[i]
				j--
				break
			}
		}

		for ;j>=i;j-- {
			if data[j] < key {
				data[i], data[j] = data[j], data[i]
				i++
				break
			}
		}
	}
	data[0], data[i-1] = data[i-1], data[0]
	if i-1+base<key{ //在前半截
		return findSmallerstBase(data[0:i], base)
	} else { // 在后半截
		return findSmallerstBase(data[i:], i+base)
	}
}

func main()  {
	tests := []struct{
		Nums []int
		Expected int
	} {
		{ []int{18,4,8,9,16,1,14,7,19,3,0,5,2,11,6} ,10 },
	}
	for _, t := range tests {
		if findSmallerst(t.Nums) !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", findSmallerst(t.Nums), t.Nums)
		}
	}
}
