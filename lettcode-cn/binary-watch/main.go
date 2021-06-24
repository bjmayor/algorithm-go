package main

import "fmt"

var digits map[int][]int
func init()  {
	digits = make(map[int][]int)
	for i:=0;i<=59;i++ {
		n:=0
		for j:=i;j>0;j=j>>1 {
			n+= j&1
		}
		digits[n]= append(digits[n], i)
	}
}

func readBinaryWatch(num int) []string {
	var result []string
	for i:=0;i<=4;i++ {
		j := num-i;
		if j > 6 || j<0 {
			continue
		}
			for _, hour := range digits[i] {
				for _, minute := range digits[j] {
					if hour>11 || minute > 59 {
						continue
					}
					result = append(result, fmt.Sprintf("%d:%02d",hour,minute))
				}
			}
	}
	return result
}

func main() {
	//fmt.Println(readBinaryWatch(1))
	//fmt.Println(readBinaryWatch(9))
	fmt.Println(readBinaryWatch(7))
}
