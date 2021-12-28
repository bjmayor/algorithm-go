package main

func numWaterBottles(numBottles int, numExchange int) int {
	return numWaterBottles2(numBottles, numExchange)
}

func numWaterBottles1(numBottles int, numExchange int) int {
	var result int
	var empty int
	for numBottles > 0{
		//喝完酒
		result += numBottles
		empty += numBottles
		//兑换一次
		numBottles = empty/numExchange
		empty -= numBottles*numExchange
	}
	return result
}

func numWaterBottles2(numBottles int, numExchange int) int {
	if numBottles>=numExchange{
		return numBottles + (numBottles-numExchange)/(numExchange-1)+1
	}
	return numBottles
}

func main()  {
	println(numWaterBottles(9,3)) // 13
	println(numWaterBottles(15,4)) // 19
	println(numWaterBottles(5,5)) // 6
	println(numWaterBottles(2,3))  // 2
}