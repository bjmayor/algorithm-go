package main

import (
	"fmt"
	"sort"
)

/*
回溯算法是一种用于解决组合、排列、子集等问题的算法思想。它通过递归探索所有可能的解决方案，并在发现某个路径不符合条件时，及时“回退”到上一步，尝试其他可能的路径。以下是回溯算法的核心思想和特点的总结：

### 回溯的核心思想

1. **递归探索**：
   - 回溯算法通过递归函数来探索所有可能的解。每次递归调用都代表了一个决策点，算法会尝试不同的选择。

2. **状态空间树**：
   - 回溯可以被视为在一个状态空间树中进行深度优先搜索。每个节点代表一个状态，边代表从一个状态到另一个状态的转移。

3. **选择与撤销**：
   - 在每个决策点，算法会做出一个选择（例如选择一个元素），然后进入下一个状态。
   - 如果发现当前选择导致的状态不符合要求，算法会撤销这个选择（即回溯），并尝试其他可能的选择。

4. **剪枝**：
   - 在某些情况下，算法可以提前判断某个路径不可能得到有效解，从而避免不必要的计算。这种提前终止的过程称为“剪枝”。

### 回溯的步骤

1. **选择**：
   - 在每个决策点，选择一个可能的选项。

2. **探索**：
   - 进入下一个状态，继续进行选择。

3. **检查**：
   - 检查当前状态是否满足问题的要求（例如，是否达到了目标状态）。

4. **撤销**：
   - 如果当前路径不满足条件，撤销上一步的选择，返回到上一个状态，尝试其他选项。

5. **终止条件**：
   - 当达到目标状态或所有选项都被尝试后，结束递归。

### 应用场景

回溯算法广泛应用于以下问题：

- **排列与组合**：生成所有可能的排列和组合。
- **子集问题**：生成给定集合的所有子集。
- **图的遍历**：如深度优先搜索（DFS）。
- **求解约束满足问题**：如数独、八皇后问题等。

### 总结

回溯算法是一种强大的解决问题的策略，通过递归和状态的选择与撤销，能够有效地探索所有可能的解。它的灵活性和适用性使其成为解决许多组合和排列问题的首选方法。通过合理的剪枝策略，回溯算法可以显著提高效率，减少不必要的计算。

代码解释
布尔数组 used：
用于跟踪每个元素是否已经被使用。这样可以避免在生成排列时选择已经使用的元素。
路径 path：
用于存储当前生成的排列。
3. 跳过重复元素：
在循环中，首先检查当前元素是否已经被使用。
然后检查当前元素是否与前一个元素相同，并且前一个元素没有被使用。如果是这种情况，则跳过当前元素，以避免生成重复的排列。
举个例子
考虑输入数组 [1, 1, 2]：
在第一次递归调用时，假设我们选择了第一个 1（used[0] = true），然后进入下一层递归。
在下一层递归中，i 可能会指向第二个 1（nums[1]）。此时，nums[1] == nums[0]，但 used[0] 是 true，所以我们可以选择第二个 1。
如果我们在第一层递归中没有选择第一个 1，而是选择了第二个 1，此时 used[1] 为 true，而 used[0] 为 false。在这种情况下，!used[i-1] 为 true，我们会跳过第二个 1，从而避免生成重复的排列。


4. 递归与回溯：
在递归调用后，进行回溯，标记当前元素为未使用，并从路径中移除最后一个元素。
*/

func permuteUnique(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)                 // 排序以便后续去重
	used := make([]bool, len(nums)) // 布尔数组，跟踪元素是否被使用
	backtrack(nums, &ans, []int{}, used)
	return ans
}

func backtrack(nums []int, ans *[][]int, path []int, used []bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 跳过已经使用的元素
		if used[i] {
			continue
		}
		// 跳过重复的元素
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		// 选择当前元素
		used[i] = true
		path = append(path, nums[i])
		backtrack(nums, ans, path, used) // 递归调用
		// 回溯
		used[i] = false
		path = path[:len(path)-1] // 移除最后一个元素
	}
}

func main() {
	fmt.Println(permuteUnique([]int{2, 2, 1, 1}))
	fmt.Println(permuteUnique([]int{0, 1, 0, 0, 9}))
}
