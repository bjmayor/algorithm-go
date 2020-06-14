package main

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	result := make([]string, 0, 0 )
	gen(n, n, "", &result)
	return result
}

func gen(left, right int, one string,  result*[]string)  {
	if left == 0 && right == 0 {
		*result = append(*result,one )
		return
	}
	if left>0 {
		gen(left-1, right,  one+"(", result)
	}
	 if right>0 && right>left{
		 gen(left, right-1, one+")", result)
	 }
}

func main()  {
	fmt.Println(generateParenthesis(3))
}