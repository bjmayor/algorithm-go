package main

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k%2 == 1 {
		return kthGrammar(n-1, (k+1)/2)
	} else {
		return 1 - kthGrammar(n-1, (k+1)/2)
	}
}

func main() {
	println(kthGrammar(1, 1)) // 0
	println(kthGrammar(2, 1)) //0
	println(kthGrammar(2, 2)) // 1
	println(kthGrammar(3, 1)) // 0
	println(kthGrammar(3, 2)) // 1
}
