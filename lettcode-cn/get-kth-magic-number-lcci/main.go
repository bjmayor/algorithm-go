package main

func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}
	v3 := []int{3}
	v5 := []int{5}
	v7 := []int{7}
	var num int
	for i := 2; i <= k; i++ {
		min := v3[0]
		if v5[0] < min {
			min = v5[0]
			if v7[0] < min {
				min = v7[0]
				v7 = v7[1:]
				v7 = append(v7, min*7)
			} else {
				v5 = v5[1:]
				v5 = append(v5, min*5)
				v7 = append(v7, min*7)
			}

		} else if v7[0] < min {
			min = v7[0]
			v7 = v7[1:]
			v7 = append(v7, min*7)
		} else {
			v3 = v3[1:]
			v3 = append(v3, 3*min)
			v5 = append(v5, 5*min)
			v7 = append(v7, 7*min)
		}
		num = min
		//fmt.Println(num)
	}
	return num
}

func main() {
	println(getKthMagicNumber(5))
}
