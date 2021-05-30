# 字典树、并查集、红黑树和AVL树、位运算

## 字典树

## 并查集

## 红黑树和AVL树

- 都是 [平衡二叉树](https://zh.wikipedia.org/wiki/%E5%B9%B3%E8%A1%A1%E6%A0%91)
- AVL树是严格的自平衡二叉树，适合 `读多写少` 的场景，如：数据库、微博等，查找、更新时间复杂度：`O(log n)`
- 红黑树是近似平衡的二叉树，适合 `写多读少` 的场景，如：map、set等

## 位运算

| 运算符 | 含义 | 解释 | 示例 |
| ----- | ---- | ---- | ---- |
| `\|` | 按位或 | 只要有一个二进制位是1结果即为1，可以简单理解为加法 | 0011 `\|` 1011 --> 1011 |
| `&` | 按位与 | 2个二进制位都为1结果才为1，只要有一个二进制位是0结果就是0 | 0011 `&` 1011 --> 0011 |
| `~` | 按位取反 | 二进制位1变0，0变1 | `~`0011 --> 1100 |
| `^` | 按位异或 | 相异为1相同为0 | 0011 `^` 1011 --> 1000 |

### XOR 按位异或

**异或：** 相同为0，不同为1。也可以用“不进位加法”来理解
`异或操作的一些特点：`

- x ^ 0 = x：  x异或0就是x本身
- x ^ x = 0：  x异或x结果是0，因为二者所有的二进制位都相同，那异或的结构就是全0
- x ^ 1s = ~x：  x异或全1，结果就是x的取反。1s是所有位都是1的意思，也就是0取反  `注意：1s = ~0`
- x ^(~x) = 1s：  首先`~x`，x取反就是把所有二进制位1变为0，0变为1，那再跟x本身异或的，肯定是所有位上都不同，那结果就是全1
- c=a^b => a^c=b, b^c=a： 交换两个数
- a^b^c = a^(b^c) = (a^b)^c：  组合

### 实战位运算要点

- 判断奇偶性：
  - 奇数：x%2 == 1 ---> (x&1) == 1，二进制位最后一位是否是1
  - 偶数：x%2 == 0 ---> (x&1) == 0，二进制位最后一位是否是0
- x/2 ---> x>>1, mid := (left+right)/2 ---> mid := (left+right)>>1
- `X = X&(X-1) 清零最低位的1` 高频
- `X & -X => 得到最低位的1` 高频
- `X & ~X = 0` 高频

## 练习题

| Title | Code | <span id="Top">Difficulty</span> | Points |
| ----- | ---- | -------------------------------- |--------|
|[208. 实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)|[208.trie.go](208.trie.go)|M|trie|
|[212. 单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)|[212.findWords.go](212.findWords.go)|H|trie、dfs+回溯|
|[191. 位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)|[191.hammingWeight.go](191.hammingWeight.go)|S|位运算|
|[231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)|[231.isPowerOfTwo.go](231.isPowerOfTwo.go)|S|位运算|
|为了解决212 Hard问题<br />刻意练几个 `dfs+回溯` 题找找感觉<br />其中岛屿问题、排列、组合是典型的回溯思想|
|[79. 单词搜索](https://leetcode-cn.com/problems/word-search/)|[79.exist.go](79.exist.go)|M|dfs+回溯|
|[剑指 Offer 38. 字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)|[offer38.permutation.go](offer38.permutation.go)|M|dfs+回溯|
|[695. 岛屿的最大面积](https://leetcode-cn.com/problems/max-area-of-island/)|[695.maxAreaOfIsland.go](695.maxAreaOfIsland.go)|M|dfs+回溯|
|[463. 岛屿的周长](https://leetcode-cn.com/problems/island-perimeter/)|[463.islandPerimeter.go](463.islandPerimeter.go)|M|dfs+回溯|
