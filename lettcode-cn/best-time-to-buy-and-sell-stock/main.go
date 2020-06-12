package main

import "fmt"

//单调栈 算法
func maxProfit(prices []int) int {
	stack := make([]int,0)
	profit := 0
	for i := 0;i<len(prices); i++ {
		if len(stack) == 0 {
			stack = append(stack, prices[i])
		} else if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
			if prices[i] - stack[0] > profit {
				profit = prices[i] - stack[0]
			}

		} else {
			top := prices[i]
			for len(stack) > 0 {
				if stack[len(stack)-1]	> top {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
			stack = append(stack,prices[i])
		}
	}
	return profit
}


//发生利润一定是 右边的值-左边的值，所以只要记录左边的最小值，就能知道产生的最大利润，然后比之前的最大利润还大，就更新。
func maxProfit2(prices []int) int {
	if len(prices) <=1 {
		return 0
	}
	low := prices[0]
	profit := 0
	for i := 1;i<len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		} else if prices[i] -low > profit {
			profit = prices[i] - low
		}
	}
	return profit
}

func main()  {
	tests := []struct{
		Nums []int
		Expected int
	} {
		//{ []int{7,1,5,3,6,4},5},
		//{ []int{7,6,4,3,1},0},
		//{ []int{},0},
		//{ []int{1},0},
		//{ []int{1,2},1},
		//{ []int{2,1,2,0,1},1},
		//{ []int{3,3},0},
		//{ []int{2,2,5},3},
		{ []int{2,1,2,1,0,1,2},2},
	}
	for _, t := range tests {
		var real = maxProfit(t.Nums)
		if real !=  t.Expected {
			fmt.Println("expected:", t.Expected, " real:", real, t.Nums)
		}
	}
}
