package main


type MyCircularDeque struct {
	data []int
	head int
	tail int
}


func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		data: make([]int, k+1, k+1),
		head:0,
		tail: 0,
	}
}


func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head= (this.head-1+len(this.data)) % len(this.data)
	this.data[this.head] = value
	return true
}


func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail= (this.tail+1+len(this.data)) % len(this.data)
	return true
}


func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head= (this.head+1+len(this.data)) % len(this.data)
	return true
}


func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail-1+len(this.data)) % len(this.data)
	return true
}


func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.head)]
}


func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.tail-1+len(this.data))%len(this.data)]
}


func (this *MyCircularDeque) IsEmpty() bool {
	return this.tail == this.head
}


func (this *MyCircularDeque) IsFull() bool {
	return  (this.tail+1 + len(this.data)) % len(this.data) == this.head
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
func main()  {
	//["MyCircularDeque","insertFront","deleteLast","getRear","getFront","getFront","deleteFront","insertFront","insertLast","insertFront","getFront","insertFront"]
	//[[4],[9],[],[],[],[],[],[6],[5],[9],[],[6]]
	obj := Constructor(4);
	println(obj.InsertFront(9)) // true
	println(obj.DeleteLast()) // tree
	println(obj.GetRear()) // true
	println(obj.GetFront()) //false
	println(obj.GetFront()) // 2
	println(obj.DeleteFront()) // true
	println(obj.InsertFront(6)) // true
	println(obj.InsertLast(5)) // true
	println(obj.InsertFront(9)) // 4
	println(obj.GetFront()) // 4
	println(obj.InsertFront(6)) // 4

}