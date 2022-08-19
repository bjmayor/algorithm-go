package main

import "fmt"

type OrderedStream struct {
	data []string
	ptr int
}


func Constructor(n int) OrderedStream {
	return OrderedStream{data: make([]string, n+1), ptr: 1}
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey] = value
	start := this.ptr
	for this.ptr<len(this.data) && this.data[this.ptr] != "" {
		this.ptr++
	}
	return this.data[start:this.ptr]
}

func main()  {
	os := Constructor(5);
	fmt.Println(os.Insert(3, "ccccc")); // 插入 (3, "ccccc")，返回 []
	fmt.Println(os.Insert(1, "aaaaa")); // 插入 (1, "aaaaa")，返回 ["aaaaa"]
	fmt.Println(os.Insert(2, "bbbbb")); // 插入 (2, "bbbbb")，返回 ["bbbbb", "ccccc"]
	fmt.Println(os.Insert(5, "eeeee")); // 插入 (5, "eeeee")，返回 []
	fmt.Println(os.Insert(4, "ddddd")); // 插入 (4, "ddddd")，返回 ["ddddd", "eeeee"]

}