package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
3272. 统计好整数的数目

给你两个 正 整数 n 和 k 。

如果一个整数 x 满足以下条件，那么它被称为 k 回文 整数 。

x 是一个 回文整数 。
x 能被 k 整除。
如果一个整数的数位重新排列后能得到一个 k 回文整数 ，那么我们称这个整数为 好 整数。比方说，k = 2 ，那么 2020 可以重新排列得到 2002 ，2002 是一个 k 回文串，所以 2020 是一个好整数。而 1010 无法重新排列数位得到一个 k 回文整数。

请你返回 n 个数位的整数中，有多少个 好 整数。

注意 ，任何整数在重新排列数位之前或者之后 都不能 有前导 0 。比方说 1010 不能重排列得到 101 。
*/

/**
 * @description 生成n位的且能被k整除的回文数
 * @param n 位数
 * @param k 整除数
 * @return []int 所有满足条件的回文数
 */
func generateKDivisiblePalindromes(n, k int) []int {
	result := []int{}

	// 计算需要生成的左半部分的位数
	leftLen := (n + 1) / 2

	// 生成左半部分的起始和结束数字
	start := 1
	for i := 1; i < leftLen; i++ {
		start *= 10
	}
	end := start * 10

	// 对每个左半部分数字，生成对应的回文数
	for i := start; i < end; i++ {
		// 获取左半部分的数字
		left := i
		// 构建右半部分
		right := 0
		temp := left
		if n%2 == 1 { // 奇数位数，需要去掉中间数字
			temp /= 10
		}
		// 反转构建右半部分
		for temp > 0 {
			right = right*10 + temp%10
			temp /= 10
		}
		// 组合成完整的回文数
		num := left
		for j := 0; j < n/2; j++ {
			num *= 10
		}
		num += right

		// 只保留能被k整除的数
		if num%k == 0 {
			result = append(result, num)
		}
	}

	return result
}

/**
 * @description 获取数字的所有数位并返回排序后的字符串key
 * @param num 输入数字
 * @return string 排序后的数位字符串
 */
func getDigitsKey(num int) string {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)

	var sb strings.Builder
	for _, d := range digits {
		sb.WriteByte(byte(d) + '0')
	}
	return sb.String()
}

/**
 * @description 计算阶乘
 * @param n 数字
 * @return int 阶乘结果
 */
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

/**
 * @description 计算数字的排列组合数
 * @param key string 排序后的数位字符串
 * @return int 可能的排列数
 */
func calculatePermutations(key string) int {
	n := len(key)
	if n == 0 {
		return 0
	}

	// 统计每个数字出现的次数
	count := make(map[byte]int)
	zeroCount := 0
	for i := 0; i < len(key); i++ {
		count[key[i]]++
		if key[i] == '0' {
			zeroCount++
		}
	}

	// 如果所有数字都相同
	if len(count) == 1 {
		return 1
	}

	// 计算分母（各个数字的阶乘的乘积）
	denominator := 1
	for _, c := range count {
		denominator *= factorial(c)
	}

	// 使用公式：(n-c0)*(n-1)!/∏ci!
	// 其中n是总位数，c0是0的个数
	return (n - zeroCount) * factorial(n-1) / denominator
}

/**
 * @description 统计好整数的数目
 * @param n 位数
 * @param k 整除数
 * @return int64 好整数的数量
 */
func countGoodIntegers(n int, k int) int64 {
	// 生成所有n位且能被k整除的回文数
	palindromes := generateKDivisiblePalindromes(n, k)

	result := 0
	seen := make(map[string]bool)

	// 检查每个回文数的排列组合
	for _, num := range palindromes {
		key := getDigitsKey(num)
		if !seen[key] {
			seen[key] = true
			result += calculatePermutations(key)
		}
	}

	return int64(result)
}

func main() {
	fmt.Println(countGoodIntegers(3, 5)) // 27
	fmt.Println(countGoodIntegers(1, 4)) // 2
	fmt.Println(countGoodIntegers(5, 6)) // 2468
	fmt.Println(countGoodIntegers(8, 5)) // test case
}
