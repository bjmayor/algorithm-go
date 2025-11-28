# DFS套路详解

## 前言
树形结构的问题是面试高频考点，而DFS（深度优先搜索）是解决树问题的核心武器。本文将系统梳理树形DFS的通用套路，并以LeetCode 2872题为例详细讲解。

## 一、树形DFS的三大核心套路

### 1. 自顶向下（Top-Down）
- **特征**：从根节点向下传递信息
- **适用场景**：路径和、深度计算、祖先信息传递
- **模板**：
```go
func dfs(node int, parentInfo int) {
    // 使用父节点传来的信息
    currentInfo := process(parentInfo)
    
    // 递归处理子节点
    for _, child := range children[node] {
        dfs(child, currentInfo)
    }
}
```

### 2. 自底向上（Bottom-Up）
- **特征**：从叶子节点向上收集信息
- **适用场景**：子树和、子树大小、子树最值
- **模板**：
```go
func dfs(node int, parent int) ReturnType {
    // 收集所有子节点的信息
    var result ReturnType
    for _, child := range children[node] {
        if child == parent continue
        childResult := dfs(child, node)
        result = combine(result, childResult)
    }
    
    // 处理当前节点
    result = processCurrentNode(node, result)
    return result
}
```

### 3. 树形DP
- **特征**：结合动态规划思想，维护多个状态
- **适用场景**：选/不选问题、最优化问题
- **模板**：
```go
func dfs(node int, parent int) (select, notSelect int) {
    // 初始化状态
    select, notSelect := value[node], 0
    
    for _, child := range children[node] {
        childSelect, childNotSelect := dfs(child, node)
        // 状态转移
        select += childNotSelect
        notSelect += max(childSelect, childNotSelect)
    }
    return
}
```

## 二、以LeetCode 2872为例详解

### 题目分析
**2872. 可以被 K 整除连通块的最大数目**

给定一棵n个节点的树，每个节点有权值value[i]，求删除若干条边后，能得到的"节点值之和能被k整除"的连通块的最大数量。

### 关键洞察
1. **贪心策略**：当某个子树的和能被k整除时，立即切断它与父节点的连接是最优的
2. **为什么贪心最优**：
   - 切断后：该子树成为一个独立的k-可整除块（+1）
   - 不切断：该子树的和向上传递，可能"污染"父节点的和，使父节点无法整除

3. **DFS类型**：典型的**自底向上**模式，需要收集子树信息

### 解题步骤

#### Step 1: 建图（邻接表）
```go
// 将边列表转换为邻接表
graph := make([][]int, n)
for _, edge := range edges {
    u, v := edge[0], edge[1]
    graph[u] = append(graph[u], v)
    graph[v] = append(graph[v], u)
}
```

#### Step 2: DFS收集子树和
```go
var ans int

func dfs(node, parent int) int {
    // 当前节点的初始和
    sum := values[node]
    
    // 收集所有子树的和
    for _, child := range graph[node] {
        if child == parent {
            continue // 避免回到父节点
        }
        childSum := dfs(child, node)
        sum += childSum
    }
    
    // 关键判断：如果当前子树和能被k整除
    if sum % k == 0 {
        ans++ // 切断与父节点的边，形成独立块
        return 0 // 向上传递0，表示这部分已经独立
    }
    
    return sum // 否则向上传递和
}
```

#### Step 3: 启动DFS
```go
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
    // 建图
    graph := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        graph[u] = append(graph[u], v)
        graph[v] = append(graph[v], u)
    }
    
    ans := 0
    dfs(0, -1) // 从节点0开始，-1表示无父节点
    return ans
}
```

### 复杂度分析
- **时间复杂度**：O(n)，每个节点访问一次
- **空间复杂度**：O(n)，递归栈深度和邻接表空间

## 三、DFS做题checklist

遇到树形问题时，按以下步骤思考：

1. **确定信息流向**
   - [ ] 需要从上往下传递信息？→ 自顶向下
   - [ ] 需要从下往上收集信息？→ 自底向上
   - [ ] 需要维护多个状态？→ 树形DP

2. **明确递归三要素**
   - [ ] 返回值：子树需要向父节点返回什么信息？
   - [ ] 参数：父节点需要向子节点传递什么信息？
   - [ ] 终止条件：叶子节点的处理逻辑

3. **实现细节**
   - [ ] 建图：边列表转邻接表
   - [ ] 防环：记录parent避免回溯
   - [ ] 状态合并：如何combine子节点的结果

4. **贪心思考**
   - [ ] 是否存在局部最优即全局最优的性质？
   - [ ] 能否在遍历过程中直接做出决策？

## 四、总结

树形DFS的本质是**信息在树上的传递和聚合**：
- 自顶向下：父节点告诉子节点"你该知道什么"
- 自底向上：子节点告诉父节点"我这里发生了什么"
- 树形DP：每个节点维护多个可能的状态

对于2872这道题，核心是理解**自底向上收集子树和 + 贪心切断**的组合策略。掌握这个套路后，大部分树形问题都能迎刃而解。

## 五、扩展练习

相似题目推荐（巩固自底向上模式）：
- LeetCode 337: 打家劫舍 III（树形DP）
- LeetCode 543: 二叉树的直径（收集子树信息）
- LeetCode 968: 监控二叉树（贪心+DFS）
- LeetCode 1617: 统计子树中城市之间最大距离（树形DFS+枚举）