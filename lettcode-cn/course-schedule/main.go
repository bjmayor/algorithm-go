package main

import "fmt"

/*
*

  - canFinish 判断是否可能完成所有课程

  - @param numCourses 课程总数

  - @param prerequisites 课程依赖关系

  - @return bool 是否可以完成所有课程

  - @complexity 时间复杂度：O(V + E)，其中 V 为课程数，E 为依赖关系数

  - @complexity 空间复杂度：O(V + E)

    示例：
    好的，让我用一个具体例子来说明这个算法的执行过程。

假设有4门课程（编号0-3），依赖关系如下：
```
prerequisites = [[1,0], [2,1], [3,2]]
```
这表示：
- 学习课程1前必须完成课程0
- 学习课程2前必须完成课程1
- 学习课程3前必须完成课程2

让我们一步步执行算法：

1. 初始化数据结构：
```
入度表inDegree:    [0, 1, 1, 1]  // 记录每个课程的前置课程数
邻接表adjacency:
0 -> [1]   // 课程0的后续课程是1
1 -> [2]   // 课程1的后续课程是2
2 -> [3]   // 课程2的后续课程是3
3 -> []    // 课程3没有后续课程
```

2. 执行过程：
```
第1步：
- 找到入度为0的课程：课程0
- 队列：[0]
- 已访问数count：0

第2步：
- 取出课程0
- count = 1
- 将课程0的后续课程1的入度减1
- 入度表变为：[0, 0, 1, 1]
- 课程1入度变为0，加入队列
- 队列：[1]

第3步：
- 取出课程1
- count = 2
- 将课程1的后续课程2的入度减1
- 入度表变为：[0, 0, 0, 1]
- 课程2入度变为0，加入队列
- 队列：[2]

第4步：
- 取出课程2
- count = 3
- 将课程2的后续课程3的入度减1
- 入度表变为：[0, 0, 0, 0]
- 课程3入度变为0，加入队列
- 队列：[3]

第5步：
- 取出课程3
- count = 4
- 课程3没有后续课程
- 队列为空，结束
```

3. 结果判断：
- count = 4 等于课程总数
- 返回 true，表示可以完成所有课程

最终的学习顺序是：0 -> 1 -> 2 -> 3

如果我们修改依赖关系，加入环：
```
prerequisites = [[1,0], [0,1]]  // 课程0和1互相依赖
```

执行过程：
1. 入度表：[1, 1]
2. 没有入度为0的课程
3. 队列为空，直接结束
4. count = 0 ≠ 课程总数
5. 返回 false，表示无法完成所有课程

让我详细讲解下 Kahn 算法。
Kahn 算法是一种用于拓扑排序的算法，其核心思想是：
找到所有入度为0的节点（没有依赖的节点）
删除这些节点以及它们的出边
重复这个过程，直到图中没有入度为0的节点

Kahn算法的特点：
可以检测图中是否存在环
如果存在环，则无法完成拓扑排序
时间复杂度为O(V+E)，其中V是节点数，E是边数
空间复杂度为O(V)，主要用于存储入度表和队列
这个算法在很多实际场景中都有应用：
课程安排
项目依赖管理
编译系统的依赖解析
任务调度系统
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建入度表和邻接表
	inDegree := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjacency[i] = make([]int, 0)
	}

	// 填充入度表和邻接表
	for _, pre := range prerequisites {
		course := pre[0]
		prerequisite := pre[1]
		inDegree[course]++
		adjacency[prerequisite] = append(adjacency[prerequisite], course)
	}

	// 将所有入度为0的课程加入队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS遍历
	count := 0
	for len(queue) > 0 {
		// 取出一个入度为0的课程
		prerequisite := queue[0]
		queue = queue[1:]
		count++

		// 将所有依赖该课程的课程的入度减1
		for _, course := range adjacency[prerequisite] {
			inDegree[course]--
			// 如果入度变为0，加入队列
			if inDegree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}

	// 如果所有课程都被访问过，说明没有环
	return count == numCourses
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites)) // 输出: true

	prerequisites = [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites)) // 输出: false
}
