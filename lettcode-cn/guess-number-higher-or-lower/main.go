package main
func guess(num int) int {
	//return 6-num
	return 1-num
}
func guessNumber(n int) int {
	left := 1
	right := n
	for left<=right {
		mid := left+(right-left)/2
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g > 0 {
			left = mid+1
		} else {
			right = mid-1
		}
	}
	return left
}
func main()  {
	println(guessNumber(1))
}

