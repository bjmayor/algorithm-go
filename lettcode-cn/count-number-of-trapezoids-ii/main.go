package main

import "fmt"

func main() {
	points := [][]int{{-73, -50}, {-35, -50}, {-15, 62}, {96, -89}, {-46, 64}, {81, -40}, {-73, -12}, {-46, -68}, {81, 14}, {-85, -9}, {39, -89}, {81, 17}, {35, -40}, {-46, -37}, {-46, -50}, {-46, -89}, {73, -40}}
	fmt.Println(countTrapezoids(points)) // 应该输出 70
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 用结构体表示斜率和截距（避免浮点精度问题）
type Fraction struct{ num, den int }

func newFraction(num, den int) Fraction {
	if den == 0 {
		return Fraction{1, 0}
	} // 表示无穷大
	if den < 0 {
		num, den = -num, -den
	}
	g := gcd(num, den)
	if g != 0 {
		num, den = num/g, den/g
	}
	return Fraction{num, den}
}

func countTrapezoids(points [][]int) int {
	n := len(points)

	// 斜率 -> 截距列表
	slopeToIntercept := make(map[Fraction][]Fraction)
	// 中点 -> 斜率列表
	midToSlope := make(map[[2]int][]Fraction)

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx, dy := x2-x1, y2-y1

			// 斜率 k = dy/dx
			k := newFraction(dy, dx)

			// 截距 b = y - kx = (y*dx - x*dy) / dx
			var b Fraction
			if dx == 0 {
				b = Fraction{x1, 1} // 垂直线用 x 坐标作为"截距"
			} else {
				b = newFraction(y1*dx-x1*dy, dx)
			}

			// 中点 (x1+x2, y1+y2)，不除以2，保持整数
			mid := [2]int{x1 + x2, y1 + y2}

			slopeToIntercept[k] = append(slopeToIntercept[k], b)
			midToSlope[mid] = append(midToSlope[mid], k)
		}
	}

	ans := 0

	// 统计有效平行线段对：同斜率、不同截距
	for _, intercepts := range slopeToIntercept {
		cnt := make(map[Fraction]int)
		for _, b := range intercepts {
			cnt[b]++
		}

		total := 0
		for _, c := range cnt {
			ans += total * c // 不同截距之间的配对
			total += c
		}
	}

	// 减去平行四边形：同中点、不同斜率
	for _, slopes := range midToSlope {
		cnt := make(map[Fraction]int)
		for _, k := range slopes {
			cnt[k]++
		}

		total := 0
		for _, c := range cnt {
			ans -= total * c // 不同斜率之间的配对
			total += c
		}
	}

	return ans
}
