package main
type pair struct{ start, end int }
type MyCalendarTwo map[int]pair

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (tree MyCalendarTwo) update(start, end, val, l, r, idx int) {
	if r < start || end < l {
		return
	}
	// 线段树查找, 第idx节点的区间是[l,r]
	// 区间无法继续拆了。直接修改
	if start <= l && r <= end {
		p := tree[idx]
		p.start += val
		p.end += val
		tree[idx] = p
		return
	}
	// 拆分区间
	mid := (l + r) >> 1
	tree.update(start, end, val, l, mid, 2*idx)
	tree.update(start, end, val, mid+1, r, 2*idx+1)
	p := tree[idx]
	// 更新父节点
	p.start = p.end + max(tree[2*idx].start, tree[2*idx+1].start)
	tree[idx] = p
}

func (tree MyCalendarTwo) Book(start, end int) bool {
	// 先加进入，start, end-1 都+1
	tree.update(start, end-1, 1, 0, 1e9, 1)
	if tree[1].start> 2 {
		// >2 表示重叠次数>2了。 不应该加入进来。 所以start, end-1, 都-1, 重置回去
		tree.update(start, end-1, -1, 0, 1e9, 1)
		return false
	}
	return true
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}



func main()  {
	  obj := Constructor();
	  println(obj.Book(10,20))// true
	println(obj.Book(50,60)) // true
	println(obj.Book(10,40)) // true
	println(obj.Book(5,15)) // false
	println(obj.Book(5,10)) // true
	println(obj.Book(25,55)) // true

}

