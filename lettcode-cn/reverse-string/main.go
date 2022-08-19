package main

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = tmp
	}
}
func main() {
	s := []byte("hello")
	reverseString(s)
	println(string(s))
}
