package main

import "fmt"

/*
*
2116. 判断一个括号字符串是否有效
一个括号字符串是只由 '(' 和 ')' 组成的 非空 字符串。如果一个字符串满足下面 任意 一个条件，那么它就是有效的：

字符串为 ().
它可以表示为 AB（A 与 B 连接），其中A 和 B 都是有效括号字符串。
它可以表示为 (A) ，其中 A 是一个有效括号字符串。
给你一个括号字符串 s 和一个字符串 locked ，两者长度都为 n 。locked 是一个二进制字符串，只包含 '0' 和 '1' 。对于 locked 中 每一个 下标 i ：

如果 locked[i] 是 '1' ，你 不能 改变 s[i] 。
如果 locked[i] 是 '0' ，你 可以 将 s[i] 变为 '(' 或者 ')' 。
如果你可以将 s 变为有效括号字符串，请你返回 true ，否则返回 false 。
*/
func canBeValid(s string, locked string) bool {
	n := len(s)
	// 长度为奇数时一定无法配对
	if n%2 == 1 {
		return false
	}

	// 从左到右检查
	balance := 0
	for i := 0; i < n; i++ {
		if locked[i] == '1' && s[i] == ')' {
			balance--
		} else {
			balance++
		}
		if balance < 0 {
			return false
		}
	}

	// 从右到左检查
	balance = 0
	for i := n - 1; i >= 0; i-- {
		if locked[i] == '1' && s[i] == '(' {
			balance--
		} else {
			balance++
		}
		if balance < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canBeValid("()))", "0101"))     // true
	fmt.Println(canBeValid("()", "11"))         // true
	fmt.Println(canBeValid("))()))", "010100")) // true
	fmt.Println(canBeValid("(())", "0000"))     // true
	fmt.Println(canBeValid(")", "0"))           // false
}
