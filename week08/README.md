# 字典树、并查集、红黑树和AVL树、位运算

## 字典树

## 并查集

## 红黑树和AVL树

- 都是 [平衡二叉树](https://zh.wikipedia.org/wiki/%E5%B9%B3%E8%A1%A1%E6%A0%91)
- AVL树是严格的自平衡二叉树，适合 `读多写少` 的场景，如：数据库、微博等，查找、更新时间复杂度：`O(log n)`
- 红黑树是近似平衡的二叉树，适合 `写多读少` 的场景，如：map、set等

## 位运算

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)|[208.trie.go](208.trie.go)|M|trie|
|[212. 单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)|[212.findWords.go](212.findWords.go)|H|trie、dfs+回溯|
|为了解决212 Hard问题<br />刻意练几个 `dfs+回溯` 题找找感觉<br />其中岛屿问题、排列、组合是典型的回溯思想|
|[79. 单词搜索](https://leetcode-cn.com/problems/word-search/)|[79.exist.go](79.exist.go)|M|dfs+回溯|
|[剑指 Offer 38. 字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)|[offer38.permutation.go](offer38.permutation.go)|M|dfs+回溯|
|[695. 岛屿的最大面积](https://leetcode-cn.com/problems/max-area-of-island/)|[695.maxAreaOfIsland.go](695.maxAreaOfIsland.go)|M|dfs+回溯|
|[463. 岛屿的周长](https://leetcode-cn.com/problems/island-perimeter/)|[463.islandPerimeter.go](463.islandPerimeter.go)|M|dfs+回溯|
