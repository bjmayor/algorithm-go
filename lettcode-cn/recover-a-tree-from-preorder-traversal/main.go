package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	if len(S) == 0 {
		return nil
	}
	m := make(map[int]*TreeNode)
	root := TreeNode{ }
	level, val, left := pickOne(S)
	root.Val = val
	m[0] = &root
	for len(left) > 0 {
		level, val, left = pickOne(left)
		parent := m[level-1]
		current := TreeNode{}
		current.Val = val
		if parent.Left == nil {
			parent.Left = &current
		} else {
			parent.Right = &current
		}
		m[level] = &current
	}

	return &root
}


func pickOne(data string) (level, val int, left string){
		level = 0
		for i:=0;i<len(data);i++ {
			if data[i] != '-' {
				j := i
				for ;j<len(data);j++ {
					if data[j] == '-' {
						break
					}
				}
				val , _ = strconv.Atoi(data[i:j])
				left = data[j:]
				return
			} else {
				level++
			}
		}
	return
}


func main()  {
	n := recoverFromPreorder("1-2--3--4-5--6--7")
	fmt.Println(n)
}