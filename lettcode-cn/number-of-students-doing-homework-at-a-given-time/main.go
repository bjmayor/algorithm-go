package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt := 0
	for i, start := range startTime {
		if start <= queryTime {
			if endTime[i] >= queryTime {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4)) // 1
}
