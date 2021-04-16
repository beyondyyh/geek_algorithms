# 泛型递归、树的递归、分治、回溯

## 名词解释

### 泛型递归

**递归通用模板模板：**

```java
public void recursion(int level, int param) {   
    // terminator 终结者，先找返回条件 
    if (level > MAX_LEVEL) {     
        // process result 处理返回结果
        return;
    }

    // process current logic 处理当前逻辑
    process(level, param);
    
    // drill down  向下进入下一层
    recursion(level: level + 1, newParam);   
    
    // restore current status 销毁当前状态
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

## 分治、回溯

## 练习题

| Title | Code | <span id="Top">Difficulty</span> |
| ----- | ---- | -------------------------------- |
|[爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)|[70.climbStairs.go](70.climbStairs.go)|S|
|[括号生成](https://leetcode-cn.com/problems/generate-parentheses/)|[22.generateParenthesis.go](22.generateParenthesis.go)|M|
|[翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/description/)|[226.invertTree.go](226.invertTree.go)|S|
|[验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)|[98.isValidBST.go](98.isValidBST.go)|M|
|[二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)|[104.maxDepth.go](104.maxDepth.go)|S|
