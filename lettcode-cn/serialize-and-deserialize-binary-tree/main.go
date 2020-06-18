package main

import (
    "fmt"
    "strconv"
)

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

type Codec struct {

}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
	    return ""
    }
    var result string
     t := make([]*TreeNode, 0)
     t = append(t, root)
     for len(t) >0 {
        top := t[0]
        t = t[1:]
        if top != nil {
            result += strconv.Itoa(top.Val)
            result += ","
            t = append(t, top.Left)
            t = append(t, top.Right)
        } else {
            result+="#,"
        }
     }
     for result[len(result)-2:] == "#," {
         result = result[:len(result)-2]
     }
     return result
}

func (this *Codec) pickOne(data string) (val, left string){
    if data[0] == '#' {
        val = data[0:1]
        left = data[2:]
    } else {
        for i:=1;i<len(data);i++ {
           if data[i] == ',' {
               val = data[0:i]
               left = data[i+1:]
               break
           }
        }

    }
    return
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
        return nil
    }
    var root TreeNode
	val, data := this.pickOne(data)
	root.Val, _ = strconv.Atoi(val)
    t := make([]*TreeNode, 0)
    t = append(t, &root)
	for len(t) >0 && len(data)>0 {
	    top := t[0]
	    t = t[1:]
        val, data = this.pickOne(data)
	    if val == "#" {
	        top.Left = nil
        } else {
            top.Left = &TreeNode{ }
            top.Left.Val, _ = strconv.Atoi(val)
            t = append(t, top.Left)
        }
        if len(data) == 0 {
            break
        }
        val, data = this.pickOne(data)
        if val == "#" {
            top.Right= nil
        } else {
            top.Right = &TreeNode{ }
            top.Right.Val, _ = strconv.Atoi(val)
            t = append(t, top.Right)
        }
    }

    return &root
}



func main()  {

    n1 := TreeNode{
        Val:   1,
        Left:  nil,
        Right: nil,
    }

    n2 := TreeNode{
        Val:   2,
        Left:  nil,
        Right: nil,
    }

    n3 := TreeNode{
        Val:   3,
        Left:  nil,
        Right: nil,
    }

    n4 := TreeNode{
       Val:   4,
       Left:  nil,
       Right: nil,
    }

    n5 := TreeNode{
        Val:   5,
        Left:  nil,
        Right: nil,
    }
    n1.Left = &n2
    n1.Right= &n3
    n3.Left= &n4
    n3.Right = &n5

     obj := Constructor();
     data := obj.serialize(&n1);
     fmt.Println(data)
    ans := obj.deserialize(data);
    fmt.Println(obj.serialize(ans))

}
