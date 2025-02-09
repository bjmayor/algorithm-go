package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 第一步：复制每个节点并插入到原链表中
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val}
		cur.Next.Next = next
		cur = next
	}

	// 第二步：设置随机指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 第三步：拆分链表
	cur = head
	dummy := &Node{Val: 0}
	copyCur := dummy
	for cur != nil {
		next := cur.Next.Next
		copyCur.Next = cur.Next
		copyCur = copyCur.Next
		cur.Next = next
		cur = next
	}

	return dummy.Next
}

func main() {
	// 示例链表构建
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Random = head.Next // 1的随机指针指向2
	head.Next.Random = head // 2的随机指针指向1

	// 复制链表
	copiedHead := copyRandomList(head)

	// 打印结果
	fmt.Println(copiedHead.Val)             // 输出 1
	fmt.Println(copiedHead.Random.Val)      // 输出 2
	fmt.Println(copiedHead.Next.Random.Val) // 输出 1
}
