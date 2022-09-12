package main

func isHappy(n int) bool {
	if n < 10 {
		return n == 1
	}
	m := map[int]struct{}{}
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
		v := 0
		for n != 0 {
			v += (n % 10) * (n % 10)
			n = n / 10
		}
		n = v
	}
	return true
}

func main() {
	println(isHappy(19))
}
