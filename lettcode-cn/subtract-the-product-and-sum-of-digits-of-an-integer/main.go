package main

func subtractProductAndSum(n int) int {
	p, q := 1, 0
	for n > 0 {
		v := n % 10
		n = n / 10
		p *= v
		q += v
	}
	return p - q
}

func main() {
	println(subtractProductAndSum(234))
	println(subtractProductAndSum(4421))
}