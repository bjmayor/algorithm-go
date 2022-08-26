package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Log struct {
	Identity int
	Flag     string
	Time     int
	Duration int
}

func parse(s string) Log {
	fields := strings.Split(s, ":")
	identity, flag, time := fields[0], fields[1], fields[2]
	i, _ := strconv.Atoi(identity)
	end, _ := strconv.Atoi(time)
	return Log{
		Identity: i,
		Flag:     flag,
		Time:     end,
	}
}
func exclusiveTime(n int, logs []string) []int {
	all := make([]*Log, 0)
	stack := make([]*Log, 0)
	times := make([]int, n)
	for _, v := range logs {
		log := parse(v)
		if log.Flag[0] == 's' {
			stack = append(stack, &log)
			all = append(all, &log)
		} else {
			top := stack[len(stack)-1]
			duration := log.Time - top.Time + 1
			top.Duration += duration
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				top.Duration -= duration
			}
		}
	}
	for _, log := range all {
		times[log.Identity] += log.Duration
	}
	return times
}

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))                         //[3,4]
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"})) //[8]
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"})) // [7,1]
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"})) // [8,1]
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"})) // [6]
}

func exclusiveTime2(n int, logs []string) []int {
	ans := make([]int, n)
	type pair struct{ idx, timestamp int }
	st := []pair{}
	for _, log := range logs {
		sp := strings.Split(log, ":")
		idx, _ := strconv.Atoi(sp[0])
		timestamp, _ := strconv.Atoi(sp[2])
		if sp[1][0] == 's' {
			if len(st) > 0 {
				ans[st[len(st)-1].idx] += timestamp - st[len(st)-1].timestamp
				st[len(st)-1].timestamp = timestamp
			}
			st = append(st, pair{idx, timestamp})
		} else {
			p := st[len(st)-1]
			st = st[:len(st)-1]
			ans[p.idx] += timestamp - p.timestamp + 1
			if len(st) > 0 {
				st[len(st)-1].timestamp = timestamp + 1
			}
		}
	}
	return ans
}
