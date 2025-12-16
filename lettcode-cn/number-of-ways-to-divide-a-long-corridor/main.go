package main

const mod = 1_000_000_007

func numberOfWays(corridor string) int {
	seatCount := 0
	plantCount := 0
	result := 1

	for _, ch := range corridor {
		if ch == 'S' {
			seatCount++
			// 当遇到第3个、第5个、第7个...座位时（奇数位且不是第1个）
			// 这时候前面的植物数量决定了分隔方案数
			if seatCount > 2 && seatCount%2 == 1 {
				result = (result * (plantCount + 1)) % mod
				plantCount = 0
			} else if seatCount%2 == 0 {
				// 遇到偶数位座位，重置植物计数
				plantCount = 0
			}
		} else {
			// 只在偶数个座位后统计植物（即在两段之间）
			if seatCount > 0 && seatCount%2 == 0 {
				plantCount++
			}
		}
	}

	// 座位数必须是2的倍数且至少有2个
	if seatCount < 2 || seatCount%2 != 0 {
		return 0
	}
	return result
}
