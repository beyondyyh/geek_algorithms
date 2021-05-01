# DFS、BFS、贪心算法、二分查找

## 名词解释

### DFS深度优先搜索

DFS包含前中后序遍历，一般都是递归写法，非递归写法使用 `栈`，或 `白灰颜色标记法`

**通用模板：**

```python
visited = set()
def DFS(node):
    # terminator, already visited
    if node in visited:
        return
    visited.add(node)
    # process current node
    # ... # logic here
    DFS(node.left) 
    DFS(node.right)
```

### BFS广度优先搜索

BFS即层级遍历，一般是用 `队列`

**通用模板：**

```python
def BFS(graph, start, end):
    queue = []
    queue.append([start])
    visited.add(start)

    while queue:
        node = queue.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        queue.push(nodes)
```

### 贪心算法

贪心算法是一种在每一步选择中都采取在当前状态下最好或最优解，从而以局部最优解推导出全局最优解的思想

**贪心算法与动态规划：**

二者都是对每个子问题的解决方案作出选择，不同点在于：

- 贪心算法不能回退
- 动态规划则会保存以前的运算结果，并`根据以前的结果对当前进行选择，有回退功能`

贪心算法一般用来解决：图中的最小生成树、求哈夫曼编码等问题。

### 二分查找

**二分查找的前提：**

- 模板函数单调性，单调递增或单调递减
- 存在左右界
- 通过通过索引访问（index accessible）

**代码模板：**

```golang
func binarySearch(nums []int, target int) bool {
    left, right := 0, len(nums)-1
    for left <= right {
        // 注意：mid = (left + right) / 2 可能存在越界，建议下面的写法
        mid := left + (right-left) / 2
        if nums[mid] == target {
            return true
        } else nums[mid] > target { // 在左半部分查找
            right = mid - 1
        } else { // 在有半部分查找
            left = mid + 1
        }
    }
    return false
}
```

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)|[102.levelOrder.go](102.levelOrder.go)|M|BFS|
|[433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)|[433.minMutation.go](433.minMutation.go)|M|DFS|
|[22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)|[22.generateParenthesis.go](22.generateParenthesis.go)|M|DFS|
|[515. 在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)|[515.largestValues.go](515.largestValues.go)|M|BFS|
|[199. 二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)|[199.rightSideView.go](199.rightSideView.go)|M|BFS|
|[127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)|[127.ladderLength.go](127.ladderLength.go)|H|DFS|
|[126. 单词接龙 II](https://leetcode-cn.com/problems/word-ladder-ii/)|[126.findLadders.go](126.findLadders.go)|H|DFS|
|[529. 扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)|[529.updateBoard.go](529.updateBoard.go)|M|DFS|
|[200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)|[200.numIslands.go](200.numIslands.go)|M|DFS|
|[322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/)|[322.coinChange.go](322.coinChange.go)|M|贪心算法|
|[860. 柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/)|[860.lemonadeChange.go](860.lemonadeChange.go)|S|贪心算法|
|[122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)|[122.maxProfit.go](122.maxProfit.go)|S|贪心算法|
|[455. 分发饼干](https://leetcode-cn.com/problems/assign-cookies/)|[455.findContentChildren.go](455.findContentChildren.go)|S|贪心算法|
|[874. 模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/)|[874.robotSim.go](874.robotSim.go)|S|贪心算法|
|[55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)|[55.canJump.go](55.canJump.go)|M|贪心算法|
|[45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)|[45.jump.go](45.jump.go)|M|贪心算法|
|[69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)|[69.mySqrt.go](69.mySqrt)|S|二分查找|
|[367. 有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)|[367.isPerfectSquare.go](367.isPerfectSquare.go)|S|二分查找|
|[33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)|[33.search.go](33.search.go)|M|二分查找|
|[74. 搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)|[74.searchMatrix.go](74.searchMatrix.go)|M|二分查找|
|[153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)|[153.findMin.go](153.findMin.go)|M|二分查找|
|[154. 寻找旋转排序数组中的最小值 II](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)|[154.findMinWithDup.go](154.findMinWithDup.go)|H|二分查找|
