package main

import (
    "fmt"
    "math"
)

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
func maxPathSum(root *TreeNode) int {
	var result  = math.MinInt32
    findmax(root, &result)
	return result
}

func findmax(root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }
    left := findmax(root.Left, result)
    right := findmax(root.Right, result)
    var max = root.Val
    if root.Val + left > max {
        max = root.Val+left
    }
    if root.Val + right > max {
        max = root.Val+right
    }
    tmp := max
    if root.Val + left + right > max {
        max = root.Val+left +right
    }
    if *result < max {
        *result = max
    }
    return tmp
}


func main()  {
    n1 := TreeNode{Val:-10}
    n2 := TreeNode{Val:9}
    n3 := TreeNode{Val:20}
    n4 := TreeNode{Val:15}
    n5 := TreeNode{Val:7}
    n1.Left = &n2
    n1.Right = &n3
    n3.Left= &n4
    n3.Right= &n5
    fmt.Println(maxPathSum(&n1))
}
