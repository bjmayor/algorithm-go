package main

// 取余法
func hammingWeight1(num uint32) int {
	n := 0
	for num > 0 {
		n += int(num % 2)
		num /= 2
	}
	return n
}
// 位移法
func hammingWeight2(num uint32) int {
	n := 0
	for num > 0 {
		n += int(num & 1)
		num = num>>1
	}
	return n
}
// 直接与最后1个1
func hammingWeight(num uint32) int {
	n := 0
	for num > 0 {
		n += 1
		num =  num & (num-1)
	}
	return n
}

func main()  {
	println(hammingWeight(00000000000000000000000000001011))
	println(hammingWeight(00000000000000000000000010000000))
	println(hammingWeight(11111111111111111111111111111101))
}