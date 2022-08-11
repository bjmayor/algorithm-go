package main


func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	m := make(map[int]int)
	m[dis(p1,p2)]++
	m[dis(p1,p3)]++
	m[dis(p1,p4)]++
	m[dis(p2,p3)]++
	m[dis(p2,p4)]++
	m[dis(p3,p4)]++
	if len(m) == 2 {
		ok := 0
		for k,v := range m {
			if v==4  || v== 2{
				ok++
			}
			if k == 0 {
				return false
			}
		}
		if ok == 2 {
			return true
		}
	}
	return false
}

func dis(pa []int, pb []int) int {
	return  (pa[0] - pb[0])*(pa[0] - pb[0]) + (pa[1] - pb[1] )*(pa[1] - pb[1] )
}

func main()  {
	println(validSquare([]int{0,0},[]int{1,1},[]int{1,0},[]int{0,1})) // true
	println(validSquare([]int{0,0},[]int{1,1},[]int{1,0},[]int{0,12})) // false
	println(validSquare([]int{1,0},[]int{-1,0},[]int{0,1},[]int{0,-1})) // true
	println(validSquare([]int{0,0},[]int{1,1},[]int{0,0},[]int{1,1})) // false
}