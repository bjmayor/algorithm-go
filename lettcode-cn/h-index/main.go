package main

import (
	"fmt"
	"sort"
)

/*

https://leetcode.cn/problems/h-index/submissions/597900033/?envType=daily-question&envId=2025-02-09

274. H 指数

给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

根据维基百科上 h 指数的定义：h 指数是指某篇论文至少被引用 h 次，且其余论文被引用的次数均不超过 h 次。


*/

/*
* hIndex 计算研究者的 h 指数
* 算法思想:
* 1. 对 citations 数组进行排序
* 2. 遍历排序后的数组,检查每个元素是否满足 h 指数的定义
* 3. 返回满足条件的最大 h 值
 */

func hIndex(citations []int) int {
	sort.Ints(citations)

	for i := 0; i < len(citations); i++ {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}
	return 0
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
}
