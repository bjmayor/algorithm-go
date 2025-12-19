# 并查集（Union-Find）套路详解

## 什么是并查集

并查集是一种用于处理**不相交集合**的数据结构，主要支持两种操作：
- **Union（合并）**：将两个集合合并为一个集合
- **Find（查找）**：判断某个元素属于哪个集合

并查集特别适合解决**连通性**问题，例如判断两个元素是否在同一个集合中。

## 并查集的基本实现

### 核心数据结构

```go
type UnionFind struct {
    parent []int  // parent[i] 表示元素 i 的父节点
    rank   []int  // rank[i] 表示以 i 为根的树的高度（用于按秩合并）
}
```

### 三个核心操作

#### 1. 初始化
每个元素初始时都是独立的集合，自己是自己的父节点。

```go
func NewUnionFind(n int) *UnionFind {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i  // 初始时自己是自己的父节点
        rank[i] = 1
    }
    return &UnionFind{parent: parent, rank: rank}
}
```

#### 2. Find（查找）- 路径压缩优化
查找元素的根节点，同时进行路径压缩，将路径上的所有节点直接连到根节点。

```go
func (uf *UnionFind) Find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.Find(uf.parent[x])  // 路径压缩
    }
    return uf.parent[x]
}
```

**路径压缩的作用**：将树的高度降低，后续操作更快。时间复杂度接近 O(1)。

#### 3. Union（合并）- 按秩合并优化
将两个集合合并，优先将高度小的树合并到高度大的树下。

```go
func (uf *UnionFind) Union(x, y int) {
    rootX := uf.Find(x)
    rootY := uf.Find(y)
    
    if rootX == rootY {
        return  // 已经在同一集合
    }
    
    // 按秩合并：将高度小的树连到高度大的树下
    if uf.rank[rootX] < uf.rank[rootY] {
        uf.parent[rootX] = rootY
    } else if uf.rank[rootX] > uf.rank[rootY] {
        uf.parent[rootY] = rootX
    } else {
        uf.parent[rootY] = rootX
        uf.rank[rootX]++
    }
}
```

**按秩合并的作用**：避免树退化成链表，保持树的平衡。

### 辅助方法

```go
// 判断两个元素是否连通
func (uf *UnionFind) IsConnected(x, y int) bool {
    return uf.Find(x) == uf.Find(y)
}

// 重置节点（某些场景需要）
func (uf *UnionFind) Reset(x int) {
    uf.parent[x] = x
    uf.rank[x] = 1
}
```

## 并查集的应用场景

### 1. 判断连通性
- 网络连接
- 朋友圈问题
- 岛屿数量

### 2. 动态连通性
- 判断图中是否存在环
- 最小生成树（Kruskal 算法）

### 3. 分组/聚类
- 账户合并
- 社交网络分析

### 4. 带时间维度的连通性（本题）
- 需要按时间顺序处理
- 可能需要回退/重置操作

## 并查集常见套路

### 套路1：基础连通性判断
```go
uf := NewUnionFind(n)
// 建立连接
uf.Union(x, y)
// 判断连通
if uf.IsConnected(a, b) {
    // 处理逻辑
}
```

### 套路2：统计集合数量
通过统计有多少个根节点来确定集合数量。
```go
count := 0
for i := 0; i < n; i++ {
    if uf.Find(i) == i {
        count++
    }
}
```

### 套路3：带权并查集
维护节点间的关系（距离、差值等）。

### 套路4：时间维度处理
- 按时间排序
- 分时间段处理
- 需要时进行回退

## 复杂度分析

使用路径压缩和按秩合并优化后：
- **单次操作时间复杂度**：O(α(n))，α 是阿克曼函数的反函数，增长极慢，实际可看作 O(1)
- **空间复杂度**：O(n)

---

## 实战案例：LeetCode 2092. 找出知晓秘密的所有专家

### 题目描述

给定 n 个专家（编号 0 到 n-1），一个会议列表 meetings[i] = [xi, yi, timei]，以及初始知道秘密的 firstPerson。

- 专家 0 在时间 0 时知道秘密，并分享给 firstPerson
- 每次会议时，知道秘密的专家会分享给对方
- 同一时间的秘密传播是瞬时的
- 返回所有最终知道秘密的专家列表

### 解题分析

这是一道**带时间维度的并查集问题**，体现了"套路4"的应用。

#### 核心难点

1. **时间维度**：需要按时间顺序处理会议
2. **同时刻处理**：同一时刻的会议必须一起处理（瞬时传播）
3. **回退机制**：如果某人没有连通到专家0，需要撤销其加入

#### 解题步骤

1. **初始化**：将专家 0 和 firstPerson 合并到同一集合
2. **按时间排序**：将所有会议按时间升序排列
3. **分时间段处理**：
   - 收集同一时刻的所有会议
   - 对这些会议进行 Union 操作
   - 检查参会者是否与专家 0 连通
   - **关键**：不连通的参会者需要 Reset（回退）
4. **收集结果**：统计所有与专家 0 连通的专家

#### 代码实现

```go
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
    uf := NewUnionFind(n)
    uf.Union(0, firstPerson)
    
    // 按时间排序
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][2] < meetings[j][2]
    })
    
    // 按时间分组处理
    i := 0
    for i < len(meetings) {
        currentTime := meetings[i][2]
        peopleInMeetings := make(map[int]bool)
        
        // 处理同一时间的所有会议
        j := i
        for j < len(meetings) && meetings[j][2] == currentTime {
            x, y := meetings[j][0], meetings[j][1]
            uf.Union(x, y)
            peopleInMeetings[x] = true
            peopleInMeetings[y] = true
            j++
        }
        
        // 回退：不连通的参会者需要重置
        for person := range peopleInMeetings {
            if !uf.IsConnected(person, 0) {
                uf.Reset(person)
            }
        }
        
        i = j
    }
    
    // 收集结果
    result := make([]int, 0)
    for i := 0; i < n; i++ {
        if uf.IsConnected(i, 0) {
            result = append(result, i)
        }
    }
    return result
}
```

### 复杂度分析

- **时间复杂度**：O(m log m + m α(n))
  - 排序：O(m log m)
  - 并查集操作：O(m α(n)) ≈ O(m)
- **空间复杂度**：O(n)

### 关键技巧总结

1. **时间分组**：同一时间的操作必须批量处理
2. **回退机制**：使用 Reset 撤销不合法的合并
3. **连通性判断**：只保留与"源点"连通的节点

## 总结

并查集是处理动态连通性问题的利器，掌握以下要点：

✅ **基本操作**：Find（路径压缩） + Union（按秩合并）  
✅ **常见套路**：连通性判断、集合统计、时间维度处理  
✅ **优化技巧**：路径压缩 + 按秩合并，时间复杂度接近 O(1)  
✅ **灵活应用**：根据题目需求添加 Reset、权重等扩展功能

通过这道题目，我们学会了如何处理**带时间维度的并查集问题**，这是并查集高级应用的典型场景。

