package main

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := heihtOf(root)
	m := height
	n := int(math.Pow(2, (float64(height))) - 1)
	var res = make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
	}
	res[0][(n-1)/2] = strconv.Itoa(root.Val)
	fillChild(root, res, 0, (n-1)/2)
	return res
}

func fillChild(node *TreeNode, res [][]string, row, col int) {
	if node.Left != nil {
		colLeft := col - int(math.Pow(2, float64(len(res)-row-2)))
		res[row+1][colLeft] = strconv.Itoa(node.Left.Val)
		fillChild(node.Left, res, row+1, colLeft)
	}
	if node.Right != nil {
		colRight := col + int(math.Pow(2, float64(len(res)-row-2)))
		res[row+1][colRight] = strconv.Itoa(node.Right.Val)
		fillChild(node.Right, res, row+1, colRight)
	}
}

func heihtOf(root *TreeNode) int {
	p := []*TreeNode{root}
	height := 0
	for len(p) > 0 {
		tmp := p
		height += 1
		for i := 0; i < len(tmp); i++ {
			n := tmp[i]
			if n.Left != nil {
				p = append(p, n.Left)
			}
			if n.Right != nil {
				p = append(p, n.Right)
			}
		}
		p = p[len(tmp):]
	}
	return height
}

func main() {
	p2 := &TreeNode{Val: 2}
	p1 := &TreeNode{
		Val:   1,
		Left:  p2,
		Right: nil,
	}

	fmt.Println(printTree(p1))
}
