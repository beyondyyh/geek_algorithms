# 高级搜索、剪枝、双向BFS、启发式搜索(A*)

## 高级搜索

### DFS

**DFS递归写法模板：**

```python
visited = set() 
def dfs(node, visited):
    if node in visited: # terminator
        # already visited
        return
    visited.add(node)
    
    # process current node here.
    ...
    for next_node in node.children():
        if next_node not in visited:
            dfs(next_node, visited)
```

**DFS非递归写法模板：**

```python
def DFS(self, tree):
    if tree.root is None:
        return []
    visited, stack = [], [tree.root]
    
    while stack:
        node = stack.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_nodes(node)
        stack.push(nodes)
        # other processing work
        ...
```

### BFS

### 双向BFS

@todo，感觉好难啊AA~~~~，还在找感觉当中，待补充个人理解

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)|[36.isValidSudoku.go](36.isValidSudoku.go)|M|双层遍历、boxIndex技巧：(row/3)*3+col/3|
|[37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/)|[37.solveSudoku.go](37.solveSudoku.go)|H|36子问题，递归、回溯|
|[51. N 皇后](https://leetcode-cn.com/problems/n-queens/)|[51.solveNQueens.go](51.solveNQueens.go)|H||
|[127. 单词接龙](https://leetcode-cn.com/problems/word-ladder/)|[127.ladderLength.go](127.ladderLength.go)|H|BFS、双向BFS|
|[433. 最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)|[433.minMutation.go](433.minMutation.go)|M|DFS+回溯、单向BFS、双向BFS|
