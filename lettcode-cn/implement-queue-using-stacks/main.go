package main

import "fmt"

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


type MyQueue struct {
	s *Stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := MyQueue{}
	q.s = NewStack()
	return q
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	var v int
	var cache []int
	for !this.s.Empty() {
		v = this.s.Pop()
		cache = append(cache, v)
	}
	this.s.Push(x)
	for i:=len(cache)-1;i>=0;i-- {
		this.s.Push(cache[i])
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.s.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.s.Peek()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s.Empty()
}

func main()  {
	  obj := Constructor();
	  obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Peek())
	 fmt.Println(obj.Pop())
}
