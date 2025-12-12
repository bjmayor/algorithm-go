package main

import (
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	userState := make([]int, numberOfUsers) // 0:online, 负数:offline,  当用户离线时，值修改为-60, 然随着时间，更新值就行，当 >=0时，修改为 0 表示重新上线。
	result := make([]int, numberOfUsers)
	for i := 0; i < numberOfUsers; i++ {
		userState[i] = 0
		result[i] = 0
	}
	// 按时间戳排序events, 可能的结构是
	// ["MESSAGE", "timestampi", "mentions_stringi"]
	// ["OFFLINE", "timestampi", "idi"]
	// 当时间戳相同时，OFFLINE事件应该先于MESSAGE事件处理
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] == events[j][1] {
			// 相同时间戳时，OFFLINE在MESSAGE之前
			return events[i][0] == "OFFLINE"
		}
		ti, _ := strconv.Atoi(events[i][1])
		tj, _ := strconv.Atoi(events[j][1])
		return ti < tj
	})
	lastTimestamp := 0
	for _, event := range events {
		timestamp, _ := strconv.Atoi(event[1])

		if timestamp > lastTimestamp {
			diff := timestamp - lastTimestamp
			lastTimestamp = timestamp
			if diff > 0 {
				for i := 0; i < numberOfUsers; i++ {
					userState[i] = min(0, userState[i]+diff)
				}
			}
		}
		if event[0] == "MESSAGE" {
			// 处理消息事件
			mentions := strings.Split(event[2], " ")
			for _, mention := range mentions {
				if mention == "ALL" {
					// 提及所有用户
					for i := 0; i < numberOfUsers; i++ {
						result[i]++
					}
				} else if mention == "HERE" {
					// 提及所有在线用户
					for i := 0; i < numberOfUsers; i++ {
						if userState[i] == 0 {
							result[i]++
						}
					}
				} else if mention != "" {
					// 提及特定用户，解析格式为 "id{数字}"
					if strings.HasPrefix(mention, "id") {
						idStr := strings.TrimPrefix(mention, "id")
						id, _ := strconv.Atoi(idStr)
						result[id]++
					}
				}
			}
		} else if event[0] == "OFFLINE" {
			// 处理离线事件
			id, _ := strconv.Atoi(event[2])
			userState[id] = -60
		}
	}
	return result
}
