package main

 type TreeNode struct {
     Val   int
     Left  *TreeNode
     Right *TreeNode
 }
 //这个最不好。因为反复计算了。不能在中间检测出问题时，马上退出。
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left != nil  && root.Left.Val >= root.Val{
        return false
    }

    if root.Right != nil  && root.Right.Val <= root.Val{
        return false
    }
    if root.Left != nil {
        _, max := findMinMaxBST(root.Left)
        if max >= root.Val {
            return false
        }
    }

    if root.Right != nil {
        min, _ := findMinMaxBST(root.Right)
        if min <= root.Val {
            return false
        }
    }

    return isValidBST(root.Left)  && isValidBST(root.Right)
}

func findMinMaxBST(root *TreeNode) (min,max int) {
    min = root.Val
    max = root.Val
    if root.Left != nil {
        t1, t2 := findMinMaxBST(root.Left)
        if t1 < min {
            min = t1
        }
        if t2 > max {
            max = t2
        }
    }
    if root.Right != nil {
        t1, t2 := findMinMaxBST(root.Right)
        if t1 < min {
            min = t1
        }
        if t2 > max {
            max = t2
        }
    }
    return
}

//中序遍历, 因为遍历出的结构一定是递增的序列，所以保存前驱，只要<=前驱，就说明不是合法的
func isValidBST1(root *TreeNode) bool {
	pre = nil
    return TraverBST(root)
}
var pre *TreeNode
func TraverBST(root *TreeNode, ) bool {
    if root == nil {
        return true
    }
    if root.Left != nil {
        if ok := TraverBST(root.Left); !ok {
            return false
        }
    }
    if pre != nil && pre.Val >= root.Val {
        return false
    }
    pre = root
    if root.Right!= nil {
        if ok := TraverBST(root.Right); !ok {
            return false
        }
    }

    return true
}

//递归解, 用lower,upper 去限定区间
func isValidBST2(root *TreeNode) bool {
    return reverse(root, nil, nil)
}

func reverse(root *TreeNode, lower *int, upper *int) bool  {
    if root == nil {
        return true
    }

    if lower != nil {
        if root.Val <= *lower {
            return false
        }
    }
    if upper != nil {
        if root.Val >= *upper {
            return false
        }
    }
    if root.Left != nil {
        if ok := reverse(root.Left, lower, &root.Val); !ok {
            return false
        }
    }

    if root.Right!= nil {
        if ok := reverse(root.Right, &root.Val, upper); !ok {
            return false
        }
    }
    return true
}