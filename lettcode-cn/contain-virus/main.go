package main

type Point struct {
	x, y int
}
var dirs = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func containVirus(isInfected [][]int) (ans int) {
	m, n := len(isInfected), len(isInfected[0])
	for {
		neighbors := []map[Point]struct{}{} // 邻居数最多的就是接下来需要 加防火墙的
		firewalls := []int{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] != 1 {
					continue
				}
				q := []Point{{i, j}}// 找到一个病毒区域，广度遍历
				neighbor := map[Point]struct{}{}
				firewall, idx := 0, len(neighbors)+1
				isInfected[i][j] = -idx // idx 是指第几个 病毒区域。1开始

				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					for _, d := range dirs {
						if x,y := p.x+d.x, p.y+d.y; 0<=x && x<m && y>=0 && y<n {
							if isInfected[x][y] == 1 {
								q = append(q, Point{x,y})
								isInfected[x][y] = -idx
							} else {
								firewall++
								neighbor[Point{i,j}] = struct{}{}
							}
						}
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}

		}
		//  全是病毒了
		if len(neighbors) == 0 {
			break //
		}

		idx :=0 // idx 是 最危险的病毒区域
		for i := 1; i < len(neighbors); i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}

		ans += firewalls[idx] // 加入了这么多防火墙
		for _, row := range isInfected {
			for j, v := range row {
				// 会继续入侵
				if v < 0 {
					if v+1 != -idx {
						row[j] = 1 // 重置为病毒状态
					} else {
						row[j] = 2 // 已经被隔离
					}
				}
			}
		}

		// 感染邻居
		for i, neighbor := range neighbors {
			if i != idx {
				for p := range neighbor {
					isInfected[p.x][p.y] = 1
				}
			}
		}

		if len(neighbors) == 1 {
			break // 只有一个病毒区域，已经被隔离
		}

	}

	return
}


func main() {
	println(containVirus([][]int{[]int{0,1,0,0,0,0,0,1},[]int{0,1,0,0,0,0,0,1},[]int{0,0,0,0,0,0,0,1},[]int{0,0,0,0,0,0,0,0}}))
}
