package main

import (
	"math/rand"
)

// maxLevel 定义跳表的最大层数
const maxLevel = 32

// pFactor 定义跳表的概率因子,用于随机层数生成
const pFactor = 0.25

/**
 * SkiplistNode 跳表节点结构
 * @property {int} val - 节点存储的值
 * @property {[]*SkiplistNode} forward - 每一层对应的下一个节点指针数组
 */
type SkiplistNode struct {
	val     int             // 存的值
	forward []*SkiplistNode // 对应每一层的下一个元素, 索引是层数
}

/**
 * Skiplist 跳表结构
 * @property {*SkiplistNode} head - 跳表头节点
 * @property {int} level - 当前跳表的层数
 */
type Skiplist struct {
	head  *SkiplistNode
	level int
}

/**
 * Constructor 构造一个新的跳表
 * @return {Skiplist} - 返回初始化的跳表结构
 */
func Constructor() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

/**
 * randomLevel 随机生成层数
 * @return {int} - 返回随机生成的层数
 */
func (Skiplist) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

/**
 * Search 在跳表中搜索目标值
 * @param {int} target - 要搜索的目标值
 * @return {bool} - 如果找到返回true,否则返回false
 */
func (s *Skiplist) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 target 的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	// 检测当前元素的值是否等于 target
	return curr != nil && curr.val == target
}

/**
 * Add 向跳表中添加一个新值
 * @param {int} num - 要添加的值
 */
func (s *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel) // 找到每一层的前置元素
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := s.randomLevel()
	s.level = max(s.level, lv)
	newNode := &SkiplistNode{num, make([]*SkiplistNode, lv)}
	// 每个层都插入了
	for i, node := range update[:lv] {
		// 对第 i 层的状态进行更新，将当前元素的 forward 指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

/**
 * Erase 从跳表中删除一个值
 * @param {int} num - 要删除的值
 * @return {bool} - 如果成功删除返回true,否则返回false
 */
func (s *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回 false
	if curr == nil || curr.val != num {
		return false
	}
	// 自底向上
	for i := 0; i < s.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的 level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}

/**
 * max 返回两个整数中的较大值
 * @param {int} a - 第一个整数
 * @param {int} b - 第二个整数
 * @return {int} - 返回较大的整数
 */
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * main 主函数,用于测试跳表的各项功能
 */
func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Add(3)
	param_1 := obj.Search(0) //false
	println(param_1)
	obj.Add(4)
	param_2 := obj.Search(1)
	println(param_2)       // true
	println(obj.Erase(0))  // false
	println(obj.Erase(1))  // true
	println(obj.Search(1)) // false
}
