package main

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum := 0
	next := 1
	for true {
		sum += next
		if sum >= target && (sum-target)%2 == 0 {
			return next
		}
		next += 1
	}
	return 0
}

func main() {
	//println(reachNumber(2)) //3
	println(reachNumber(3)) // 2
}
