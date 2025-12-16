# 练习指南

## 📋 需要补全的代码

已为你移除了核心实现，保留了框架和提示。需要补全以下部分：

### 1️⃣ 构建邻接表 (第13-14行)
```go
// TODO: 将 hierarchy 转换为邻接表
// 注意：节点编号从1开始，需要-1转为0索引
```

### 2️⃣ 初始化子节点收益数组 (第24-26行)
```go
// TODO: 创建 subProfit0 和 subProfit1 数组
```

### 3️⃣ 分组背包合并子节点 (第32-37行)
```go
// TODO: 
// 1. 累加 childResult.size 到 uSize
// 2. 倒序遍历预算 (从 budget 到 0)
// 3. 内层循环到 min(childResult.size, i)
// 4. 更新 subProfit0[i] 和 subProfit1[i]
```

### 4️⃣ 计算当前节点的DP值 (第41-50行)
```go
// TODO:
// dp0[i] = max(不买当前节点, 买当前节点原价)
// dp1[i] = max(不买当前节点, 买当前节点半价)
```

## 💡 提示

### 分组背包关键点
- **倒序遍历**：避免重复使用子节点
- **size 剪枝**：子树最多花费 size，超过无意义
- **状态转移**：`newProfit[i] = max(oldProfit[i], oldProfit[i-sub] + childProfit[sub])`

### DP状态转移
- **dp0[i]**：父节点没买，当前节点可以原价买或不买
- **dp1[i]**：父节点买了，当前节点可以半价买或不买
- 买节点时：`profit = future[u] - cost + subProfit1[剩余预算]`

## 🧪 测试

补全后运行：
```bash
go test -v
```

所有测试通过即完成！查看完整解决方案：`main_solution.go`
