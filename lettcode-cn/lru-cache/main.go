package main

import "fmt"

type LinkNode struct {
	key int
	value int
	next *LinkNode
	pre *LinkNode
}
type LRUCache struct {
	head *LinkNode
	tail *LinkNode
	capacity int
	len int
	m map[int]*LinkNode//缓存下，以达到O(1)的效率
}


func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.head = nil
	cache.capacity = capacity
	cache.m = make(map[int]*LinkNode)
	return cache
}


func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; !ok {
		return -1
	} else {
		if n == this.head {
			return this.head.value
		}
		n.pre.next = n.next
		if n.next!=nil {
			n.next.pre = n.pre
		} else {
			this.tail = n.pre
		}
		this.head.pre,n.next,n.pre, this.head = n, this.head, nil,n

		return this.head.value
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if this.Get(key) == -1 {
		n := &LinkNode{
			key:   key,
			value: value,
		}
		if this.head != nil {
			this.head, n.next, this.head.pre = n, this.head, n
		} else {
			this.head = n
			this.tail = n
		}
		this.m[key] = n
		this.len++
		if this.len > this.capacity {
			delete(this.m, this.tail.key)
			this.tail.pre.next, this.tail = nil, this.tail.pre
			this.len--
		}
	} else {
		this.head.value = value
	}
}

func main()  {
	cache := Constructor(2)

	cache.Put(1, 1);
	cache.Put(2, 2);
	fmt.Println(cache.Get(1))      // 返回  1
	cache.Put(3, 3);    // 该操作会使得关键字 2 作废
	fmt.Println(cache.Get(2));       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得关键字 1 作废
	fmt.Println(cache.Get(1));       // 返回 -1 (未找到)
	fmt.Println(cache.Get(3));       // 返回  3
	fmt.Println(cache.Get(4));       // 返回  4

	//cache := Constructor(1)
	//
	//cache.Put(2, 1);
	//fmt.Println(cache.Get(2))      // 返回  1
}
