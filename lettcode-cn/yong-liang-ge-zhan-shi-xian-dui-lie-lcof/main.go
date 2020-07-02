package main


type Stack struct {
	data []int
}

func NewStack() *Stack {
	s := Stack{}
	s.data = make([]int,0,0)
	return &s
}
func (this *Stack) Push(x int)  {
	this.data = append(this.data, x)
}

func (this *Stack) Pop() int {
	v := this.Peek()
	this.data = this.data[:len(this.data)-1]
	return v
}

func (this *Stack) Empty() bool  {
	return len(this.data) == 0
}


func (this *Stack) Peek()  int{
	return this.data[len(this.data)-1]
}


type CQueue struct {
 head Stack
 tail Stack
}


func Constructor() CQueue {
	q := CQueue{}
	q.head = Stack{}
	q.tail = Stack{}
	return q
}


func (this *CQueue) AppendTail(value int)  {
	this.tail.Push(value)
}


func (this *CQueue) DeleteHead() int {
	if !this.head.Empty() {
		return this.head.Pop()
	}
	for !this.tail.Empty() {
		this.head.Push(this.tail.Pop())
	}
	if !this.head.Empty() {
		return this.head.Pop()
	}
	return -1
}

func main()  {
	obj := Constructor();
	obj.AppendTail(3);
	obj.DeleteHead()
	obj.DeleteHead()
}


