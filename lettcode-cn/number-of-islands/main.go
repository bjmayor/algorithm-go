package main

import "fmt"

type UnionF struct {
	data []int
}

func NewUnionF(n int) *UnionF {
	u := UnionF{}
	u.data = make([]int, n, n)
	for i:=0;i<n;i++ {
		u.data[i] = i
	}
	return &u
}

func (u *UnionF)FindRoot(b int) int {
	root := b
	for u.data[root]	!= root {
			root = u.data[root]
	}
	tmp := b
	for u.data[tmp]	!= root {
		tmp,  u.data[tmp] = u.data[tmp],root
	}
	return root
}

func (u *UnionF)remove(b int) {
	u.data[b] = -1
}

func (u *UnionF)Union(x,y int)  {
	rootx := u.FindRoot(x)
	rooty := u.FindRoot(y)
	if rootx != rooty {
		u.data[rootx] = rooty
	}
}

func (u *UnionF)count() int {
	cnt := 0
	for i:=0;i<len(u.data);i++ {
		if u.data[i] == i {
			cnt++
		}
	}
	return cnt
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	u := NewUnionF(len(grid)*len(grid[0]))
	for i:=0;i<len(grid);i++ {
		for j:=0;j<len(grid[i]);j++ {
			if grid[i][j] == '1' {
				if i-1 >=0 && grid[i-1][j] == '1' {
					u.Union((i-1)*len(grid[i])+j, i*len(grid[i])+j)
				}
				if j-1>=0 && grid[i][j-1] == '1' {
					u.Union(i*len(grid[i])+j-1, i*len(grid[i])+j, )
				}
			} else {
				u.remove(i*len(grid[i])+j)
			}
		}
	}
	return u.count()
}

func main()  {
	////3
	fmt.Println(numIslands([][]byte{
		{'1','1','0','0','0'},
		{'1','1','0','0','0'},
		{'0','0','1','0','0'},
		{'0','0','0','1','1'},
	}))
	//
	////1
	fmt.Println(numIslands([][]byte{
		{'1','1','1'},
		{'0','1','0'},
		{'1','1','1'},
	}))

	//1
	fmt.Println(numIslands([][]byte{
		{'1','0','1','1','1'},
		{'1','0','1','0','1'},
		{'1','1','1','0','1'},
	}))
}