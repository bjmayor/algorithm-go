package main
type MyCircularQueue struct {
  front int // 头部元素
  rear int // 尾部元素
  data []int
}


func Constructor(k int) MyCircularQueue {
  return MyCircularQueue{
    front: 0,
    rear: 0,
    data: make([]int, k+1,k+1), // 多一个用于区分 满和空
  }
}

func (this *MyCircularQueue) index(value int)int {
  return (value+len(this.data))%len(this.data)
}

func (this *MyCircularQueue) EnQueue(value int) bool {
  if this.IsFull() {
    return false
  }
  this.data[this.index(this.rear)] = value
  this.rear = this.index(this.rear+1)
  return true
}


func (this *MyCircularQueue) DeQueue() bool {
  if this.IsEmpty() {
    return false
  }
  this.front = this.index(this.front+1)
  return true
}


func (this *MyCircularQueue) Front() int {
  if this.IsEmpty() {
    return -1
  }
  return this.data[this.index(this.front)]
}


func (this *MyCircularQueue) Rear() int {
  if this.IsEmpty() {
    return -1
  }
  return this.data[this.index(this.rear-1)]
}


func (this *MyCircularQueue) IsEmpty() bool {
  return this.rear == this.front
}


func (this *MyCircularQueue) IsFull() bool {
  return (this.rear+1) % len(this.data) == this.front
}

func main()  {
  //circularQueue := Constructor(3); // 设置长度为 3
  //println(circularQueue.EnQueue(1)); // 返回 true
  //println(circularQueue.EnQueue(2)); // 返回 true
  //println(circularQueue.EnQueue(3)); // 返回 true
  //println(circularQueue.EnQueue(4)); // 返回 false，队列已满
  //println(circularQueue.Rear()); // 返回 3
  //println(circularQueue.IsFull()); // 返回 true
  //println(circularQueue.DeQueue()); // 返回 true
  //println(circularQueue.EnQueue(4)); // 返回 true
  //println(circularQueue.Rear()); // 返回 4
  circularQueue := Constructor(3); // 设置长度为 3
  println(circularQueue.EnQueue(1)); // 返回 true
  println(circularQueue.EnQueue(2)); // 返回 true
  println(circularQueue.EnQueue(3)); // 返回 true
  println(circularQueue.EnQueue(4)); // 返回 true
  println(circularQueue.Rear()); // 返回 3
  println(circularQueue.DeQueue()); // 返回 true
  println(circularQueue.DeQueue()); // 返回 3
  println(circularQueue.EnQueue(4)); // 返回 true
  println(circularQueue.Rear()); // 返回 3
}