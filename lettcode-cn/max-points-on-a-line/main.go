package main

func maxPoints(points [][]int) int {
	result := 1
	for i:=0;i<len(points);i++ {
		for j:=i+1;j<len(points);j++ {
			tmp := 2
			dy := points[j][1]- points[i][1]
			dx :=  points[j][0]- points[i][0]
			for k:=j+1;k<len(points);k++ {
				dy2 := points[k][1]- points[j][1]
				dx2 :=  points[k][0]- points[j][0]
				if dy * dx2 == dx *dy2 {
					tmp++
				}
			}
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

func main()  {
	println(maxPoints([][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}}))
}