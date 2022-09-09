package main

func average(salary []int) float64 {
	low := salary[0]
	high := salary[0]
	sum := 0
	for _, v := range salary {
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
		sum += v
	}
	return float64(sum-low-high) / float64(len(salary)-2)
}

func main() {
	println(average([]int{4000, 3000, 1000, 2000}))
}
