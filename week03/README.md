# 泛型递归、树的递归、分治、回溯

## 名词解释

### 泛型递归

**递归通用模板：**

```golang
func recursion(level int, param interface{}) {   
    // terminator 终结者，先找返回条件 
    if level > MAX_LEVEL {     
        // process result 处理返回结果
        return
    }

    // process current logic 处理当前逻辑
    process(level, param);
    
    // drill down  向下进入下一层
    recursion(level + 1, newParam);   
    
    // revert current status 销毁当前状态
}
```

**递归四部曲：**

- 第一部分：递归终结者，写递归一定要先把递归终止条件写上，否则一不注意可能死循环或无限递归。
- 第二部分：处理当前逻辑块；
- 第三部分：下探到下一层，level参数用来标记当前是哪一层；
- 第四部分：清理当前层的状态，如果需要的话。

**递归误区/注意点：**

- 抵制人肉进行递归（最大误区）
- `找到最近最简单的方法，将其拆解成可重复解决的问题（找最近重复子问题）`
- 数学归纳法思维

### 树的递归

树的求解一般都是递归，因为树的定义本身就有其不断重复性

### 分治, divide_conquer

**分治代码模板：**

```c++
int divide_conquer(Problem *problem, int params) {  
    // recursion terminator  
    if (problem == nullptr) { 
        process_result    
        return return_result;  
    }

    // process current problem  
    subproblems = split_problem(problem, data)

    // drill down 下探到下一层
    subresult1 = divide_conquer(subproblem[0], p1)
    subresult2 = divide_conquer(subproblem[1], p1)
    subresult3 = divide_conquer(subproblem[2], p1)
    ...
    
    // merge  
    result = process_result(subresult1, subresult2, subresult3)  
    // revert the current level status

    return 0;
}
```

**分治四部曲：**

- terminator, 递归终止条件
- process(split your big problem)
- drill down(subproblems), merge(subresult),
- revert states

可以看出分治的本质也是递归，与泛型递归不同的地方在于：

- 分治是把大问题拆解成一个个的子问题，终结条件是没有子问题可拆解了
- 处理当前层逻辑，需要把子问题的结果 `merge在一起之后再 drill down`

> 最经典的例子就是 `归并排序`

### `回溯，Backtrack`

**使用场景：**

回溯法（backtrack）常用于遍历列表所有子集，是 DFS 深度搜索一种，一般用于全排列，穷尽所有可能，遍历的过程实际上是一个决策树的遍历过程。时间复杂度一般 O(N!)，它不像动态规划存在重叠子问题可以优化，回溯算法就是 `纯暴力穷举`，复杂度一般都很高。

**通用模板：**

```txt
result = []
func backtrack(选择列表, 路径) {
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(选择列表, 路径)
        撤销选择，这步贼拉重要
}
```

> 核心就是从选择列表里做一个选择，然后一直递归往下搜索答案，如果遇到路径不通，就返回来撤销这次选择。

**大白话自己的理解：**

- 有很多步 `选择` 要做，我们不知道每步选择应该选啥，所以，我们先随便选一个，然后前往下一步，继续做选择；
- 当我们走到某个位置，发现无路可走时，我们就往回退（return），退到上一步，重新做选择。

**典型例题：**

- [78. 子集](https://leetcode-cn.com/problems/subsets/)
- [90. 子集II](https://leetcode-cn.com/problems/subsets-ii/)
- [46. 全排列](https://leetcode-cn.com/problems/permutations/)
- [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)
- [39. 组合总和](https://leetcode-cn.com/problems/combination-sum/)
- [131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)
- [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)
- [93. 复原 IP 地址](https://leetcode-cn.com/problems/restore-ip-addresses/)
- [51. N 皇后](https://leetcode-cn.com/problems/n-queens/)

> 分治、回溯本质上是递归，找最近重复性；动态规划：找最优重复性；递归：找最近重复性

#### 回溯与动态规划的区别

**共同点：**
用于求解多阶段决策问题。多阶段决策问题即：

- 求解一个问题分为很多步骤（阶段）；
- 每一个步骤（阶段）可以有多种选择。

**不同点：**

- 动态规划只需要求我们评估最优解是多少，最优解对应的具体解是什么并不要求。因此很适合应用于评估一个方案的效果；
- 回溯算法可以搜索得到所有的方案（当然包括最优解），但是本质上它是一种遍历算法，时间复杂度很高。

## 练习题

| Title | Code | <span id="Top">Difficulty</span> |
| ----- | ---- | -------------------------------- |
|[爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)|[70.climbStairs.go](70.climbStairs.go)|S|
|[括号生成](https://leetcode-cn.com/problems/generate-parentheses/)|[22.generateParenthesis.go](22.generateParenthesis.go)|M|
|[翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/description/)|[226.invertTree.go](226.invertTree.go)|S|
|[验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)|[98.isValidBST.go](98.isValidBST.go)|M|
|[二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|[104.maxDepth.go](104.maxDepth.go)|S|
|[Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)|[50.myPow.go](50.myPow.go)|M|
|[子集](https://leetcode-cn.com/problems/subsets/)|[78.subsets.go](78.subsets.go)|M|
|[电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)|[17.letterCombinations.go](17.letterCombinations.go)|M|
|[N 皇后](https://leetcode-cn.com/problems/n-queens/)|[51.solveNQueens.go](51.solveNQueens.go)|H|
|强化|-|-|
|[从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)|[105.buildTree.go](105.buildTree.go)|M|
|[子集II](https://leetcode-cn.com/problems/subsets-ii/)|[90.subsetsWithDup.go](../week04/90.subsetsWithDup.go)|M|
|[组合](https://leetcode-cn.com/problems/combinations/)|[77.combine.go](../week04/77.combine.go)|M|
|[组合总和](https://leetcode-cn.com/problems/combination-sum/)|[39.combinationSum.go](../week04/39.combinationSum.go)|M|
|[全排列](https://leetcode-cn.com/problems/permutations/)|[46.permute.go](../week04/46.permute.go)|M|
|[全排列 II](https://leetcode-cn.com/problems/permutations-ii/)|[47.permuteUnique.go](../week04/47.permuteUnique.go)|M|
|[复原 IP 地址](https://leetcode-cn.com/problems/restore-ip-addresses/)|[93.restoreIpAddresses.go](../week04/93.restoreIpAddresses.go)|M|
