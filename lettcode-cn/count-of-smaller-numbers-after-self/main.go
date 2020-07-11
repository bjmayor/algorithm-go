package main

import (
	"fmt"
	"math"
	"sort"
)
/**
** 方法1， 暴力解  952ms, 4.3M  11行代码
** 方法2， 插入排序  82ms, 5.2M  24行代码
** 方法3， 树状数组+离散化  8ms, 4.9M 52行代码
** 方法4， 索引数组+归并排序  8ms, 4.7M 70行代码
** 方法5， 线段树  12ms, 5.4M  50行代码
** 方法6， 二叉搜索树, 8ms, 5.1M 30行代码, 综合性能和实现难度， 二叉搜索树取胜。
 */

//直接暴力法 是n平方 复杂度---952ms,4.3M
func countSmaller1(nums []int) []int {
	var res = make([]int, len(nums))
	for i:=0;i<len(nums);i++ {
		for j:=i+1;j<len(nums);j++ {
			if nums[j] < nums[i] {
				res[i]+=1
			}
		}
	}
	return res
}

//插入排序 是n平方 复杂度 ---82ms, 5.2M
func countSmaller2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	var res = make([]int, len(nums))
	var tmp []int
	tmp = append(tmp, nums[len(nums)-1])
	res[len(nums)-1] = 0
	for i:=len(nums)-2;i>=0;i-- {
		k := nums[i]
		tmp = append(tmp, k)
		j:= len(tmp)-2
		for ;j>=0;j-- {
			if tmp[j] >= k {
				tmp[j+1] = tmp[j]
			}	else {
				break
			}
		}
		tmp[j+1] = k
		res[i] = j+1
	}
	return res
}


var a, c []int //a是去重并排完序的桶, c是树状数组
//方法3 树状数组+离散化
func countSmaller3(nums []int) []int {
	resultList :=make([]int, len(nums))
	discretization(nums)
	c = make([]int, len(nums) + 5)
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(nums[i])
		resultList[i] =  query(id - 1)
		update(id)
	}
	return resultList
}

//获取c[x]管理的元素个数
func lowBit(x int) int {
	return x & (-x)
}

//pos的桶+1， 需要更新管理pos的所有节点
func update(pos int) {
	for pos < len(c)	{
		c[pos] += 1
		pos += lowBit(pos)
	}
}

func query(pos int) int {
	ret :=0
	for pos > 0 {
		ret += c[pos]
		pos -= lowBit(pos)
	}
	return ret
}

func discretization(nums []int) {
	set := map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}
	a = make([]int, 0, len(nums))
	for num := range set {
		a = append(a, num)
	}
	sort.Ints(a)
}

//返回桶的id
func getId(x int) int {
	return sort.SearchInts(a, x) + 1
}


var(
	//临时数组
	tmp   []int
	//存放虚拟索引, 存放的索引，类似指针
	index []int
	//存放结果值
	count []int
)

//方法4 归并排序实现, 8ms, 4.7M
func countSmaller4(nums []int) []int {
	if nums==nil || len(nums)==0{
		return nil
	}
	tmp=make([]int,len(nums))
	count=make([]int,len(nums))
	//存放虚拟索引
	index=make([]int,len(nums))
	//不改变nums的值，通过虚拟索引进行归并排序
	for i:=0;i<len(nums);i++{
		index[i]=i//value为数字在nums数组中的索引。index为排序以后的sorted数组中的位置。排完序后index[0] = 3
	}
	mergesort(nums,0,len(nums)-1)
	return count
}

func mergesort(nums []int, left int, right int)  {
	if left >= right {
		return
	}
	mid := (left+right)/2

	mergesort(nums, left, mid)
	mergesort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int)  {
	l := left
	r := mid+1
	k := left //tmp数组的当前索引

	for l<=mid && r <=right {
		if nums[index[l]] <= nums[index[r]] {
			tmp[k]	= index[l]
			count[index[l]] += r-(mid+1)
			l++
			k++
		} else {
			tmp[k] = index[r]
			k++
			r++
		}
	}
	for l<=mid {
		tmp[k]	= index[l]
		count[index[l]] += r-(mid+1)
		l++
		k++
	}
	for r<=right {
		tmp[k]	= index[r]
		r++
		k++
	}
	//把tmp 复制到 index
	for i:=left;i<=right;i++ {
		index[i] = tmp[i]
	}
}



type Line struct {
	start int
	end int
	mid int
	cnt int // 落在这个区间的数量
	left *Line
	right *Line
}

func NewLine(start, end int) *Line  {
	line := Line{
		start: start,
		end:   end,
		mid:   (start+end)/2,
	}
	//考虑负数
	if line.mid == line.end {
		line.mid--
	}
	return &line
}

//返回<num的元素个数
func (this *Line)add(num int) int {
	this.cnt ++
	if this.start == this.end { //元线段，不能再切分
		return 0
	}
	if this.left == nil {
		this.left = NewLine(this.start, this.mid)
	}
	if this.right == nil {
		this.right = NewLine(this.mid+1, this.end)
	}
	if num <= this.mid {
		return this.left.add(num)
	} else {
		return this.left.cnt+this.right.add(num)
	}
}

//方法5 线段树 ---- 12ms, 5.4M
func countSmaller5(nums []int) []int {
	var res = make([]int, len(nums))
	min := math.MaxInt32
	max := math.MinInt32
	for i:=0;i<len(nums);i++ {
		if nums[i]	 > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	root := NewLine(min,max)
	for i:=len(nums)-1;i>=0;i--{
		res[i] = root.add(nums[i])
	}
	return res
}

type Tree struct {
	v int
	n int //左子树节点的个数
	left *Tree
	right *Tree
}

//方法6 二叉搜索树, 复杂度 nlgn--- 8ms, 5.1M
func countSmaller6(nums []int) []int {
	var res = make([]int, len(nums))
	var root *Tree
	for i:=len(nums)-1;i>=0;i-- {
		root = insertAndcountTree(root, &Tree{v:nums[i]}, res, i)
	}
	return res
}

func insertAndcountTree(root, tree *Tree, res []int, i int)  *Tree {
	if root == nil {
		root = tree
		return root
	}
	if root.v >= tree.v { //新的节点在左边。子节点数+1
		root.n ++
		root.left = insertAndcountTree(root.left, tree, res, i)
	} else {
		res[i] += 1+root.n
		root.right = insertAndcountTree(root.right, tree, res, i)
	}
	return root
}

func main()  {
	fmt.Println(countSmaller3([]int{5,2,6,1})) //2,1,1,0
	fmt.Println(countSmaller3([]int{1,0,2})) //1,0,0
	fmt.Println(countSmaller3([]int{5,2,6666666,1})) //2,1,1,0
	fmt.Println(countSmaller3([]int{-1,-2})) //1,0
	fmt.Println(countSmaller3([]int{2,0,1})) //2,0,0
}