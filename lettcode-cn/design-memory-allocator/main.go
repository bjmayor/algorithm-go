package main

import "fmt"

/**
 * 设计内存分配器
 * https://leetcode.cn/problems/design-memory-allocator/
 *
 * 给定一个大小为 n 的内存数组,实现内存分配和释放功能
 * - memory[i] = 0 表示第 i 个内存单元未被分配
 * - memory[i] = mID 表示第 i 个内存单元已被 mID 标识的用户分配
 */

type Allocator struct {
	memory []int
}

func Constructor(n int) Allocator {
	return Allocator{
		memory: make([]int, n),
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	// 寻找连续的size个空闲内存单元
	for i := 0; i <= len(this.memory)-size; i++ {
		canAllocate := true
		// 检查从i开始的size个内存单元是否都是空闲的
		for j := 0; j < size; j++ {
			if this.memory[i+j] != 0 {
				canAllocate = false
				i = i + j // 优化:直接跳过已分配的内存
				break
			}
		}
		if canAllocate {
			// 分配内存
			for j := 0; j < size; j++ {
				this.memory[i+j] = mID
			}
			return i
		}
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	count := 0
	for i := 0; i < len(this.memory); i++ {
		if this.memory[i] == mID {
			this.memory[i] = 0
			count++
		}
	}
	return count
}

func main() {
	allocator := Constructor(10)
	fmt.Println(allocator)
	allocator.Allocate(3, 1)
	fmt.Println(allocator)
	allocator.Allocate(5, 2)
	fmt.Println(allocator)
	allocator.FreeMemory(2)
	fmt.Println(allocator)
	allocator.Allocate(6, 3)
	fmt.Println(allocator)
}
