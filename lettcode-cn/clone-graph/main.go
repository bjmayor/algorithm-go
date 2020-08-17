package main

  type Node struct {
      Val int
      Neighbors []*Node
  }

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    m := make(map[int]*Node)
    return cloneGraphHelper(node, m)
}

func cloneGraphHelper(node *Node, m map[int]*Node) *Node {
    if e, ok := m[node.Val];ok {
        return e
    }
    var ret *Node
    ret = &Node{
        Val:       node.Val,
        Neighbors: nil,
    }
    m[node.Val] = ret
    for _, n := range node.Neighbors {
        ret.Neighbors = append(ret.Neighbors, cloneGraphHelper(n,m))
    }
    return ret
}
