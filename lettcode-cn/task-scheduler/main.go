package main

func leastInterval(tasks []byte, n int) int {
	cnt := map[byte]int{}
	for _, task := range tasks {
		cnt[task]++
	}
	maxExec, maxExecCnt := 0, 0 // 重复任务的最大值，重复任务的最大值有几个
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)

}

func main() {
	println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))                               // 8
	println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0))                               // 6
	println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)) // 16
}
