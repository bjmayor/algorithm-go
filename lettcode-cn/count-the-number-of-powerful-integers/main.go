package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
2999. 统计强大整数的数目

给你三个整数 start ，finish 和 limit 。同时给你一个下标从 0 开始的字符串 s ，表示一个 正 整数。

如果一个 正 整数 x 末尾部分是 s （换句话说，s 是 x 的 后缀），且 x 中的每个数位至多是 limit ，那么我们称 x 是 强大的 。

请你返回区间 [start..finish] 内强大整数的 总数目 。

如果一个字符串 x 是 y 中某个下标开始（包括 0 ），到下标为 y.length - 1 结束的子字符串，那么我们称 x 是 y 的一个后缀。比方说，25 是 5125 的一个后缀，但不是 512 的后缀。

提示：
1 <= start <= finish <= 10^15
1 <= limit <= 9
1 <= s.length <= floor(log10(finish)) + 1
s 数位中每个数字都小于等于 limit 。
s 不包含任何前导 0 。
*/
func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	start_ := strconv.FormatInt(start-1, 10)
	finish_ := strconv.FormatInt(finish, 10)
	return calculate(finish_, s, limit) - calculate(start_, s, limit)
}

func calculate(x string, s string, limit int) int64 {
	if len(x) < len(s) {
		return 0
	}
	if len(x) == len(s) {
		if x >= s {
			return 1
		}
		return 0
	}

	suffix := x[len(x)-len(s):]
	var count int64
	preLen := len(x) - len(s)

	for i := 0; i < preLen; i++ {
		digit := int(x[i] - '0')
		if limit < digit {
			count += int64(math.Pow(float64(limit+1), float64(preLen-i)))
			return count
		}
		count += int64(digit) * int64(math.Pow(float64(limit+1), float64(preLen-1-i)))
	}
	if suffix >= s {
		count++
	}
	return count
}

func main() {
	fmt.Println(numberOfPowerfulInt(1, 6000, 4, "124"))     // 5
	fmt.Println(numberOfPowerfulInt(15, 215, 6, "10"))      // 2
	fmt.Println(numberOfPowerfulInt(1000, 2000, 4, "3000")) // 0
	fmt.Println(numberOfPowerfulInt(20, 1159, 5, "20"))     // 8
}
