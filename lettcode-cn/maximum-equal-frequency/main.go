package main

func maxEqualFreq(nums []int) (ans int) {
	freq := map[int]int{}
	count := map[int]int{}
	maxFreq := 0
	for i, num := range nums {
		if count[num] > 0 {
			freq[count[num]]--
		}
		count[num]++
		maxFreq = max(maxFreq, count[num])
		freq[count[num]]++
		if maxFreq == 1 ||
			freq[maxFreq]*maxFreq+freq[maxFreq-1]*(maxFreq-1) == i+1 && freq[maxFreq] == 1 ||
			freq[maxFreq]*maxFreq+1 == i+1 && freq[1] == 1 {
			ans = max(ans, i+1)
		}
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	//println(maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}))                      // 7
	//println(maxEqualFreq([]int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}))         // 8
	//println(maxEqualFreq([]int{1, 1}))                                        // 2
	println(maxEqualFreq([]int{1, 2, 3, 1, 2, 3, 4, 4, 4, 4, 1, 2, 3, 5, 6})) // 13
}
