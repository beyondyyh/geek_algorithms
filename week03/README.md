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

树的求解一般都是递归

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
- 处理当前层逻辑，需要把子问题的结果merge在一起之后再 drill down

> 最经典的例子就是 `归并排序`

### 回溯，Backtracking

**回溯的百科定义：**
回溯法采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。
`回溯法通常用最简单的递归方法来实现`，在反复重复上述的步骤后可能出现两种 情况:

- 找到一个可能存在的正确的答案;
- 在尝试了所有可能的分步方法后宣告该问题没有答案。

**高频题目：**

- [实现pow(x, n)即计算x的n次幂函数](https://leetcode-cn.com/problems/powx-n/)
- [求数组所有可能的子集（幂集）](https://leetcode-cn.com/problems/subsets/)

> 分治、回溯本质上是递归，找最近重复性；动态规划：找最优重复性；递归：找最近重复性

## 练习题

| Title | Code | <span id="Top">Difficulty</span> |
| ----- | ---- | -------------------------------- |
|[爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)|[70.climbStairs.go](70.climbStairs.go)|S|
|[括号生成](https://leetcode-cn.com/problems/generate-parentheses/)|[22.generateParenthesis.go](22.generateParenthesis.go)|M|
|[翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/description/)|[226.invertTree.go](226.invertTree.go)|S|
|[验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)|[98.isValidBST.go](98.isValidBST.go)|M|
|[二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|[104.maxDepth.go](104.maxDepth.go)|S|
|[Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)|[50.myPow.go](50.myPow.go)|M|
[子集](https://leetcode-cn.com/problems/subsets/)|[78.subsets.go](78.subsets.go)|M|
