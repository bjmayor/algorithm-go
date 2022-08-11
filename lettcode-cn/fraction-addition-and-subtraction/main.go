package main

import "strconv"

func fractionAddition(expression string) string {
	numerators :=[]int{}
	denominator :=[]int{}
	sign := 1
	nums := []byte{}
	for i:=0;i<len(expression);i++ {
		ch := expression[i]
		switch ch {
		case '-':
			sign = -1
			if len(nums) > 0 {
				num, _ := strconv.Atoi(string(nums))
				denominator = append(denominator,num)
			}
			nums = []byte{}
		case '+':
				sign = 1
				if len(nums) > 0 {
					num, _ := strconv.Atoi(string(nums))
					denominator = append(denominator,num)
				}
			nums = []byte{}
		case '/':
			if len(nums) > 0 {
				num, _ := strconv.Atoi(string(nums))
				numerators = append(numerators,sign * num)
			}
			nums = []byte{}
		default:
			nums = append(nums, ch)
		}

	}
	if len(nums) > 0 {
		num, _ := strconv.Atoi(string(nums))
		denominator = append(denominator,num)
	}
	leastcommonmulti := denominator[0]
	for _, v := range denominator[1:] {
		leastcommonmulti = lcm(leastcommonmulti, v)
	}
	result := 0
	for i, v := range numerators {
		result += leastcommonmulti/denominator[i]	* v
	}
	if result == 0 {
		return "0/1"
	}
	gmd := result*leastcommonmulti/lcm(result, leastcommonmulti)
	return 	strconv.Itoa(result/gmd)+"/"+strconv.Itoa(leastcommonmulti/gmd)
}

func lcm(a,b int) (i int)  {
	for i=a;;i+=a {
		if i%b==0 {
			return i;
		}
	}
}

func main()  {
	println(fractionAddition("-1/2+1/2")) // 0/1
	println(fractionAddition("-1/2+1/2+1/3")) // 1/3
	println(fractionAddition("1/3-1/2")) // -1/6
}