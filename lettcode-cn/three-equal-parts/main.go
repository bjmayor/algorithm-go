package main

func threeEqualParts(arr []int) []int {
	sum := 0
	for _, v := range arr {
		if v == 1 {
			sum++
		}
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	size := sum / 3

	first, second, third, cur := 0, 0, 0, 0
	for i, v := range arr {
		if v == 1 {
			if cur == 0 {
				first = i
			} else if cur == size {
				second = i
			} else if cur == size*2 {
				third = i
			}
			cur++
		}
	}

	n := len(arr)
	length := n - third // 第3部分不管前导0之后 的长度

	if first+length <= second && second+length <= third {
		i := 0
		// 去掉前导0后，对应的位置值必须相等
		for third+i < n {
			if arr[first+i] != arr[second+i] || arr[second+i] != arr[third+i] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + length - 1, second + length}
	}
	return []int{-1, -1}
}
