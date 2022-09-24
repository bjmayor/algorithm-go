package main

type LinkItem struct {
	Val  int
	Next *LinkItem
}
type MyLinkedList struct {
	head *LinkItem
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	n := this.head
	i := 0
	for n != nil {
		if i == index {
			return n.Val
		}
		i++
		n = n.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := &LinkItem{
		Val:  val,
		Next: this.head,
	}
	this.head = head
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := this.head
	for n != nil && n.Next != nil {
		n = n.Next
	}
	node := &LinkItem{
		Val:  val,
		Next: nil,
	}
	if n == nil {
		this.head = node
	} else {
		n.Next = node
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	n := this.head
	var p *LinkItem
	i := 0
	for n != nil {
		if i == index {
			node := &LinkItem{
				Val:  val,
				Next: n,
			}
			p.Next = node
			return
		}
		i++
		p = n
		n = n.Next
	}
	if index == i {
		this.AddAtTail(val)
		return
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	n := this.head
	var p *LinkItem
	i := 0
	for n != nil {
		if i == index {
			if p == nil {
				this.head = n.Next
				return
			}
			p.Next = n.Next
			return
		}
		i++
		p = n
		n = n.Next
	}
}

func main() {
	obj := Constructor()
	//obj.AddAtHead(1)
	//obj.AddAtTail(3)
	obj.AddAtIndex(0, 10)
	obj.AddAtIndex(0, 20)
	obj.AddAtIndex(1, 30)
	obj.Get(0)
	//obj.DeleteAtIndex(0)
	//obj.Get(1)
}
