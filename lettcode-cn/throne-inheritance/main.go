package main

import "fmt"

type ThroneInheritance struct {
	Parent *ThroneInheritance
	Name string
	Children []*ThroneInheritance
	IsDead bool
	all map[string]*ThroneInheritance
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{Name: kingName,all:make(map[string]*ThroneInheritance)}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	node := this.findNodeByName(parentName)
	child := Constructor(childName)
	child.Parent = node
	this.all[childName] = &child
	node.Children = append(node.Children, &child)
}

func (this *ThroneInheritance) findNodeByName(name string) *ThroneInheritance{
	if this.Name== name {
		return this
	}
	return this.all[name]
}

func (this *ThroneInheritance) Death(name string) {
	node := this.findNodeByName(name)
	node.IsDead = true
}

// 先序遍历递归版本
func (this *ThroneInheritance) GetInheritanceOrder1() []string {
	result := []string{}
	n := this
	if !n.IsDead {
		result = append(result, n.Name)
	}
	for _, v := range n.Children {
		result = append(result, v.GetInheritanceOrder()...)
	}
	return result
}

// 先序遍历循环版本
func (this *ThroneInheritance) GetInheritanceOrder() []string {
	result := []string{}
	n := this
	stack := []*ThroneInheritance{}
	stack = append(stack,n)
	for len(stack) > 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !t.IsDead {
			result = append(result, t.Name)
		}
		for i:=len(t.Children)-1;i>=0;i-- {
			stack = append(stack, t.Children[i])
		}
	}
	return result
}

// 这是按题意做的
func (this *ThroneInheritance) Successor(x *ThroneInheritance, set map[string]bool) *ThroneInheritance{
	found := false
	var largerChild *ThroneInheritance
	for _, c := range x.Children {
		if !set[c.Name]{
			largerChild = c
			found = true
			break
		}
	}
	if !found {
		if x.Parent == nil {
			return nil
		}	else {
			return this.Successor(x.Parent,set)
		}
	}
	return largerChild
}

func main() {
	obj := Constructor("king")
	obj.Birth("king", "Alice")
	obj.Birth("king", "Bob")
	obj.Birth("Alice", "Jack")
	fmt.Println(obj.GetInheritanceOrder())
}
