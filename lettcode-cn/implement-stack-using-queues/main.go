package main

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	s := Queue{}
	s.data = make([]int,0,0)
	return &s
}
func (this *Queue) Push(x int)  {
	this.data = append(this.data, x)
}

func (this *Queue) Pop() int {
	v := this.Peek()
	this.data = this.data[1:]
	return v
}

func (this *Queue) Empty() bool  {
	return len(this.data) == 0
}


func (this *Queue) Peek()  int{
	return this.data[0]
}


type MyStack struct {
	q *Queue
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	s :=  MyStack{}
	s.q = NewQueue()
	return s
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	var cache []int
	for !this.q.Empty() {
		v := this.q.Pop()
		cache = append(cache, v)
	}
	this.q.Push(x)
	for i:=0;i<len(cache);i++ {
		this.q.Push(cache[i])
	}
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.q.Pop()
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q.Peek()
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.Empty()
}

func main()  {
	obj := Constructor();
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	fmt.Println(obj.Pop())
}
