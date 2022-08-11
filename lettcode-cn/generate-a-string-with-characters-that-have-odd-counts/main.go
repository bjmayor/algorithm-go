package main

func generateTheString(n int) string {
	chars := "abcdefghijklmnopqrstuvwzyz"
	ret := []byte{}
	if n%2==1 {
		for i:=0;i<n;i++ {
			ret = append(ret, chars[0])
		}
	} else {
		ret = append(ret, chars[0])
		for i:=1;i<n;i++ {
			ret = append(ret, chars[1])
		}
	}
	return string(ret)
}

func main()  {
	println(generateTheString(4))
	println(generateTheString(2))
	println(generateTheString(7))
}