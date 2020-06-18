package main

import "fmt"

func maxScoreSightseeingPair(A []int) int {
	return maxScoreSightseeingPair2(A)
}


//会超时
func maxScoreSightseeingPair1(A []int) int {
	max := 0
	for i:=0;i<len(A);i++ {
		for j:=i+1;j<len(A);j++ {
			if (A[i]+A[j]+i-j) > max {
				max = 	A[i]+A[j]+i-j
			}
		}
	}
	return max
}

//看成A[i]+i  和  A[j]-j。 假如前i个元素已经得出结果为result, 当时的A[i]+i 为maxi,
// 那么当增加一个新元素j时，新的和为maxi+A[j]-j,此时和之前的result比较，更新result即可。
func maxScoreSightseeingPair2(A []int) int {
	maxi := A[0]
	var result int
	for i := 1; i < len(A); i++ {
		result = max(result, maxi+A[i]-i)
		maxi = max(maxi, A[i]+i)
	}
	return result
}

func max(x,y int) int  {
	if x > y {
		return x
	}
	return y
}


func main()  {
	fmt.Println(maxScoreSightseeingPair([]int{8,1,5,2,6}))
}
