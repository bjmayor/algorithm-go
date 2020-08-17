package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var result = ""
	var padding = ""
	for j:=len(num2)-1;j>=0;j-- {
		c := num2[j]
		one := multiplyWithChar(num1, c)
		one+=padding
		padding+="0"
		result = add(result, one)
	}
	return result
}


func multiplyWithChar(num1 string, num2 uint8) string {
	var result  = make([]uint8, len(num1)+1)
	var jin uint8
	idx := len(num1)
	for i:=len(num1)-1;i>=0;i-- {
		tmp := (num2-'0') * (num1[i]-'0')
		tmp += jin
		jin = 0
		if tmp >= 10 {
			jin = tmp/10
			tmp = tmp%10
		}
		result[idx] = tmp+'0'
		idx--
	}
	result[0] = jin+'0'
	if result[0] =='0' {
		result = result[1:]
	}
	return string(result)
}



func add(num1 string, num2 string) string {
	size := len(num1)
	if len(num2) > size {
		size = len(num2)
	}
	size++
	var result  = make([]uint8, size)
	var jin uint8
	idx := size-1
	for i:=0;i<size;i++ {
		tmp :=jin
		jin = 0
		if i<len(num1) {
			tmp += num1[len(num1)-i-1]-'0'
		}
		if i<len(num2) {
			tmp += num2[len(num2)-i-1]-'0'
		}

		if tmp >= 10 {
			jin = tmp/10
			tmp = tmp%10
		}
		result[idx] = tmp+'0'
		idx--
	}
	if result[0] =='0' {
		result = result[1:]
	}
	return string(result)
}

func main()  {
	//fmt.Println(multiply("234","978"))	//228852
	fmt.Println(multiply("123","456"))	//"56088"
}