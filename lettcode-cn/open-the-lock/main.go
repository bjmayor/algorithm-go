package main

import (
	"math"
	"strconv"
)

// 最短路径。
func openLock1(deadends []string, target string) int {
	var m  [10000]int // 已经求出的解
	var final [10000]bool// 是否已经求出解
	for _, s := range deadends {
		d, _ := strconv.Atoi(s)
		final[d] = true
	}
	if final[0] {
		return -1
	}
	final[0] = true
	for i:=1;i<=9999;i++ {
		m[i] = math.MaxInt32
	}
	for _, v := range linked(0) {
		m[v] = 1
	}

	d, _ := strconv.Atoi(target)
	if final[d] {
		return m[d]
	}
	for i:=1;i<=9999;i++ {
		min := math.MaxInt32
		// 从所有节点中找到离v0最近的点
		var k int // k为当前离v0最近的点
		for j:=0;j<=9999;j++ {
			if !final[j] && m[j]<min {
				min = m[j]
				k = j
			}
		}
		final[k] = true
		if k == d {
			return min
		}
		// 和 k 相连的点只有8个。试试这个8个点即可
		for _, v := range linked(k) {
			if !final[v] && min+1 < m[v] {
				m[v] = min+1
			}
		}

	}

	if m[d] == math.MaxInt32 {
		return -1
	}
	return m[d]
}

func linked(k int) [8]int {
	var result [8]int
	times := 10
	left := k%times
	right := k/times
	for i:=0;i<4;i++  {
		result[i*2] =  right*times + (left+times/10)%times
		result[i*2+1] =  right*times + (times+left-times/10)%times
		times *= 10
		left = k%times
		right = k/times
	}
	return result
}

// 广度优先搜索
func openLock(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[start] {
		return -1
	}

	// 枚举 status 通过一次旋转得到的数字
	get := func(status string) (ret []string) {
		s := []byte(status)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ret = append(ret, string(s))
			s[i] = b
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] && !dead[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}


func main()  {
	println(openLock([]string{"0201","0101","0102","1212","2002"}, "0202"))
	println(openLock([]string{"8888"}, "0009"))
	println(openLock([]string{"8887","8889","8878","8898","8788","8988","7888","9888"}, "8888"))
	println(openLock([]string{"0201","0101","0102","1212","2002"}, "0000"))
}